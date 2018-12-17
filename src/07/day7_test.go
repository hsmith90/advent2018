package day7

import (
	"fmt"
	"helper"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestPartA(t *testing.T) {
	lines, _ := helper.ReadLines("input")

	var stepList []*Step

	for _, line := range lines {
		regex := regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin")
		names := regex.FindStringSubmatch(line)
		mainName := names[1]
		childName := names[2]

		mainFound, mainStep := FindStep(mainName, stepList)
		childFound, childStep := FindStep(childName, stepList)

		mainStep.childSteps = append(mainStep.childSteps, childStep)
		childStep.parentSteps = append(childStep.parentSteps, mainStep)

		if !mainFound {
			mainStep.name = mainName
			time := []rune(mainName)
			mainStep.time = int(time[0]) - 4
			stepList = append(stepList, mainStep)
		}
		if !childFound {
			childStep.name = childName
			time := []rune(childName)
			childStep.time = int(time[0]) - 4
			stepList = append(stepList, childStep)
		}
	}

	var baseSteps []*Step
	for i, s := range stepList {
		if len(s.parentSteps) == 0 {
			baseSteps = append(baseSteps, stepList[i])
		}
	}

	var strBuild strings.Builder
	workerBaseSteps := make([]*Step, len(baseSteps))
	copy(workerBaseSteps, baseSteps)

	DoSteps(&strBuild, baseSteps)
	fmt.Printf("Part A: %v\n", strBuild.String())

	totalTime := DoStepsWorkers(5, workerBaseSteps, len(stepList))

	fmt.Printf("Part B: %v\n", totalTime)

}

type Step struct {
	name        string
	parentSteps []*Step
	childSteps  []*Step
	time        int
}

func FindStep(name string, list []*Step) (bool, *Step) {
	for i, item := range list {
		if name == item.name {
			return true, list[i]
		}
	}

	return false, &Step{}
}

func DoSteps(b *strings.Builder, currentChildren []*Step) {
	sort.Slice(currentChildren, func(i, j int) bool {
		return currentChildren[i].name < currentChildren[j].name
	})

	if len(currentChildren) == 0 {
		return
	}
	s := currentChildren[0]
	currentChildren = append(currentChildren[:0], currentChildren[1:]...)
	b.WriteString(s.name)

	for i, c := range s.childSteps {
		allParentsDone := true
		for _, p := range c.parentSteps {
			if !strings.Contains(b.String(), p.name) {
				allParentsDone = false
			}
		}
		if allParentsDone {
			currentChildren = append(currentChildren, s.childSteps[i])
		}
	}

	DoSteps(b, currentChildren)
}

var done []*Step
var doNext []*Step

func DoStepsWorkers(workerCount int, steps []*Step, totalSteps int) (timePassed int) {
	doNext = steps
	var workers []*Worker
	for i := workerCount; i > 0; i-- {
		workers = append(workers, &Worker{job: &Step{}, time: 0})
	}
	timePassed = -1
	fmt.Println("Sec\tW1\tW2\tW3\tW4\tW5")
	for len(done) != totalSteps {
		sort.Slice(doNext, func(i, j int) bool {
			return doNext[i].name < doNext[j].name
		})

		for i, w := range workers {
			if w.time != 0 {
				w.time--
				if w.time == 0 {
					FinishJob(w)
					StartJob(w)
					if i > 0 {
						for j := i; j >= 0; j-- {
							if workers[j].time == 0 {
								StartJob(workers[j])
							}
						}
					}
				}
			} else {
				StartJob(w)
			}
		}

		timePassed++

		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", timePassed, workers[0].job.name, workers[1].job.name, workers[2].job.name, workers[3].job.name, workers[4].job.name)
	}

	return timePassed
}

type Worker struct {
	job  *Step
	time int
}

func Contains(s *Step, list []*Step) bool {
	for _, l := range list {
		if s.name == l.name {
			return true
		}
	}
	return false
}

func FinishJob(w *Worker) {
	done = append(done, w.job)

	for i, c := range w.job.childSteps {
		allParentsDone := true
		for _, p := range c.parentSteps {
			if !Contains(p, done) {
				allParentsDone = false
			}
		}
		if allParentsDone {
			doNext = append(doNext, w.job.childSteps[i])
			sort.Slice(doNext, func(i, j int) bool {
				return doNext[i].name < doNext[j].name
			})
		}
	}
	w.job = &Step{name: "."}
}

func StartJob(w *Worker) {
	if len(doNext) > 0 {
		s := doNext[0]
		doNext = append(doNext[:0], doNext[1:]...)

		w.job = s
		w.time = s.time
	}
}
