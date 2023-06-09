package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func (p *PlanData) GetJson(path string) (*PlanData, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		err := fmt.Errorf("error openning data json file. %w", err)
		return &PlanData{}, err
	}

	var response PlanData
	err = json.Unmarshal(file, &response)
	if err != nil {
		err := fmt.Errorf("error unmarshalling json from file. %w", err)
		return &PlanData{}, err
	}
	return &response, nil
}

// FIXME: Return a single PlanData object with high priority
func (p *PlanData) PriorityFilter() PlanData {
	filteredData := removeLowPriority(*p)
	return filteredData
}

func (p *PlanData) AdjustPlanDate() {
	for i, d := range p.Plans {
		d.Schedule.StartDate = time.Now().Add(time.Hour * 24 * 30)
		p.Plans[i].Schedule.StartDate = d.Schedule.StartDate
	}
}

/* func removeLowPriority(p PlanData) PlanData {
	var aux PlanData
	filter := p
	aux.Device = filter.Device

	j := 1
	for _, v := range p.Plans {
		if v.Name == filter.Plans[j].Name {
			if v.Region.Priority > filter.Plans[j].Region.Priority {
				aux.Plans = append(aux.Plans, filter.Plans[j])
			} else {
				aux.Plans = append(aux.Plans, v)
			}
		} else {
			aux.Plans = append(aux.Plans, filter.Plans[j])
		}
		j++
		if j > (len(p.Plans) - 1) {
			break
		}
	}
	k := 1
	for i, v := range aux.Plans {
		if v.Name == aux.Plans[k].Name {
			if v.Region.Priority < filter.Plans[k].Region.Priority {
				aux.Plans = append(aux.Plans[:i], aux.Plans[i+1:]...)
			}
		}
	}

	return aux
} */

func removeLowPriority(p PlanData) PlanData {
	var aux PlanData
	filter := p
	aux.Device = filter.Device

	for i := 0; i <= len(filter.Plans)-1; i++ {
		curr := filter.Plans[i]
		for j := i + 1; j <= len(filter.Plans)-1; j++ {
			next := filter.Plans[j]
			if curr.Name == next.Name {
				if curr.Region.Priority > next.Region.Priority {
					aux.Plans = append(aux.Plans, next)
				} else if curr.Region.Priority < next.Region.Priority {
					aux.Plans = append(aux.Plans, curr)
				} else {
					j++
				}
			} else {
				i = j
				break
			}

		}
	}
	aux = removeDupes(aux)
	return aux
}

func removeDupes(p PlanData) PlanData {
	aux := p
	j := 1
	for i := 0; i <= len(p.Plans)-1; i++ {
		if p.Plans[i].Name == p.Plans[j].Name {
			aux.Plans = append(aux.Plans[:i], aux.Plans[i+1:]...)
		}
		j++
		if j <= len(p.Plans)-1 {
			break
		}
	}
	return aux
}
