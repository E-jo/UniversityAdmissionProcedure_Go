package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Applicant struct {
	FirstName, LastName string
	ExamScores          []float32
	Priorities          []string
}

const debug = false

func main() {
	// n is max students per department
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	//file, err := os.Open("C:\\Users\\erics\\Downloads\\applicant_list1.txt")
	file, err := os.Open("applicants.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var applicants []Applicant

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 10 {
			fmt.Println("Invalid input format")
			return
		}

		firstName := fields[0]
		lastName := fields[1]
		examScoreStrings := fields[2:7]
		examScores := make([]float32, len(examScoreStrings))
		for i, str := range examScoreStrings {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting string '%s' to int: %v\n", str, err)
				return
			}
			examScores[i] = float32(num)
		}
		priorities := fields[7:10]

		applicants = append(applicants, Applicant{
			FirstName:  firstName,
			LastName:   lastName,
			ExamScores: examScores,
			Priorities: priorities,
		})
	}

	var biotechApplicants []Applicant
	var chemistryApplicants []Applicant
	var engineeringApplicants []Applicant
	var mathematicsApplicants []Applicant
	var physicsApplicants []Applicant
	var meanExamScoreI, meanExamScoreJ float32

	for i := 1; i <= 3; i++ {
		// biotech
		if len(biotechApplicants) < n {
			newApplicants := filterApplicants(applicants, "Biotech", i)

			if debug {
				fmt.Printf("Biotech applicants available in round %v:\n", i)
				for _, value := range newApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[1])
				}
				fmt.Println()
			}

			// Sort new Biotech applicants by chemistry and physics exams (descending order), ties broken by full name
			sort.Slice(newApplicants, func(i, j int) bool {
				meanExamScoreI = (newApplicants[i].ExamScores[0] + newApplicants[i].ExamScores[1]) / 2
				meanExamScoreJ = (newApplicants[j].ExamScores[0] + newApplicants[j].ExamScores[1]) / 2
				deptExamScoreI := newApplicants[i].ExamScores[4]
				deptExamScoreJ := newApplicants[j].ExamScores[4]
				examScoreI := max(meanExamScoreI, deptExamScoreI)
				examScoreJ := max(meanExamScoreJ, deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := newApplicants[i].FirstName + " " + newApplicants[i].LastName
				fullNameJ := newApplicants[j].FirstName + " " + newApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			biotechApplicants = append(biotechApplicants, newApplicants...)

			if len(biotechApplicants) > n {
				biotechApplicants = biotechApplicants[:n]
			}

			// Sort Biotech applicants by chemistry and physics exams (descending order), ties broken by full name
			sort.Slice(biotechApplicants, func(i, j int) bool {
				meanExamScoreI = (biotechApplicants[i].ExamScores[0] + biotechApplicants[i].ExamScores[1]) / 2
				meanExamScoreJ = (biotechApplicants[j].ExamScores[0] + biotechApplicants[j].ExamScores[1]) / 2
				deptExamScoreI := biotechApplicants[i].ExamScores[4]
				deptExamScoreJ := biotechApplicants[j].ExamScores[4]
				examScoreI := max(meanExamScoreI, deptExamScoreI)
				examScoreJ := max(meanExamScoreJ, deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := biotechApplicants[i].FirstName + " " + biotechApplicants[i].LastName
				fullNameJ := biotechApplicants[j].FirstName + " " + biotechApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			if debug {
				fmt.Printf("Biotech applicants selected in round %v:\n", i)
				for _, value := range biotechApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[1])
				}
				fmt.Println()
			}

			// remove selected Biotech applicants from main applicant list
			applicants = removeApplicants(applicants, biotechApplicants)
		}

		// chemistry
		if len(chemistryApplicants) < n {
			newApplicants := filterApplicants(applicants, "Chemistry", i)

			if debug {
				fmt.Printf("Chemistry applicants available in round %v:\n", i)
				for _, value := range newApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[1])
				}
				fmt.Println()
			}

			// Sort new Chemistry applicants by exam (descending order), ties broken by full name
			sort.Slice(newApplicants, func(i, j int) bool {
				deptExamScoreI := newApplicants[i].ExamScores[4]
				deptExamScoreJ := newApplicants[j].ExamScores[4]
				examScoreI := max(newApplicants[i].ExamScores[1], deptExamScoreI)
				examScoreJ := max(newApplicants[j].ExamScores[1], deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := newApplicants[i].FirstName + " " + newApplicants[i].LastName
				fullNameJ := newApplicants[j].FirstName + " " + newApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			chemistryApplicants = append(chemistryApplicants, newApplicants...)

			if len(chemistryApplicants) > n {
				chemistryApplicants = chemistryApplicants[:n]
			}

			// Sort Chemistry applicants by exam (descending order), ties broken by full name
			sort.Slice(chemistryApplicants, func(i, j int) bool {
				deptExamScoreI := chemistryApplicants[i].ExamScores[4]
				deptExamScoreJ := chemistryApplicants[j].ExamScores[4]
				examScoreI := max(chemistryApplicants[i].ExamScores[1], deptExamScoreI)
				examScoreJ := max(chemistryApplicants[j].ExamScores[1], deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := chemistryApplicants[i].FirstName + " " + chemistryApplicants[i].LastName
				fullNameJ := chemistryApplicants[j].FirstName + " " + chemistryApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			if debug {
				fmt.Printf("Chemistry applicants selected in round %v:\n", i)
				for _, value := range chemistryApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[1])
				}
				fmt.Println()
			}

			// remove selected Chemistry applicants from main applicant list
			applicants = removeApplicants(applicants, chemistryApplicants)
		}

		// engineering
		if len(engineeringApplicants) < n {
			newApplicants := filterApplicants(applicants, "Engineering", i)

			if debug {
				fmt.Printf("Engineering applicants available in round %v:\n", i)
				for _, value := range newApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[3])
				}
				fmt.Println()
			}

			// Sort new Engineering applicants by exam (descending order), ties broken by full name
			sort.Slice(newApplicants, func(i, j int) bool {
				meanExamScoreI = (newApplicants[i].ExamScores[2] + newApplicants[i].ExamScores[3]) / 2
				meanExamScoreJ = (newApplicants[j].ExamScores[2] + newApplicants[j].ExamScores[3]) / 2
				deptExamScoreI := newApplicants[i].ExamScores[4]
				deptExamScoreJ := newApplicants[j].ExamScores[4]
				examScoreI := max(meanExamScoreI, deptExamScoreI)
				examScoreJ := max(meanExamScoreJ, deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := newApplicants[i].FirstName + " " + newApplicants[i].LastName
				fullNameJ := newApplicants[j].FirstName + " " + newApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			engineeringApplicants = append(engineeringApplicants, newApplicants...)

			if len(engineeringApplicants) > n {
				engineeringApplicants = engineeringApplicants[:n]
			}

			// Sort Engineering applicants by exam (descending order), ties broken by full name
			sort.Slice(engineeringApplicants, func(i, j int) bool {
				meanExamScoreI = (engineeringApplicants[i].ExamScores[2] + engineeringApplicants[i].ExamScores[3]) / 2
				meanExamScoreJ = (engineeringApplicants[j].ExamScores[2] + engineeringApplicants[j].ExamScores[3]) / 2
				deptExamScoreI := engineeringApplicants[i].ExamScores[4]
				deptExamScoreJ := engineeringApplicants[j].ExamScores[4]
				examScoreI := max(meanExamScoreI, deptExamScoreI)
				examScoreJ := max(meanExamScoreJ, deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := engineeringApplicants[i].FirstName + " " + engineeringApplicants[i].LastName
				fullNameJ := engineeringApplicants[j].FirstName + " " + engineeringApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			if debug {
				fmt.Printf("Engineering applicants selected in round %v:\n", i)
				for _, value := range engineeringApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[3])
				}
				fmt.Println()
			}

			// remove selected Engineering applicants from main applicant list
			applicants = removeApplicants(applicants, engineeringApplicants)
		}

		// mathematics
		if len(mathematicsApplicants) < n {
			newApplicants := filterApplicants(applicants, "Mathematics", i)

			if debug {
				fmt.Printf("Mathematics applicants available in round %v:\n", i)
				for _, value := range newApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[2])
				}
				fmt.Println()
			}

			// Sort new Mathematics applicants by exam (descending order), ties broken by full name
			sort.Slice(newApplicants, func(i, j int) bool {
				deptExamScoreI := newApplicants[i].ExamScores[4]
				deptExamScoreJ := newApplicants[j].ExamScores[4]
				examScoreI := max(newApplicants[i].ExamScores[2], deptExamScoreI)
				examScoreJ := max(newApplicants[j].ExamScores[2], deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := newApplicants[i].FirstName + " " + newApplicants[i].LastName
				fullNameJ := newApplicants[j].FirstName + " " + newApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			mathematicsApplicants = append(mathematicsApplicants, newApplicants...)

			if len(mathematicsApplicants) > n {
				mathematicsApplicants = mathematicsApplicants[:n]
			}

			// Sort Mathematics applicants by exam (descending order), ties broken by full name
			sort.Slice(mathematicsApplicants, func(i, j int) bool {
				deptExamScoreI := mathematicsApplicants[i].ExamScores[4]
				deptExamScoreJ := mathematicsApplicants[j].ExamScores[4]
				examScoreI := max(mathematicsApplicants[i].ExamScores[2], deptExamScoreI)
				examScoreJ := max(mathematicsApplicants[j].ExamScores[2], deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := mathematicsApplicants[i].FirstName + " " + mathematicsApplicants[i].LastName
				fullNameJ := mathematicsApplicants[j].FirstName + " " + mathematicsApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			if debug {
				fmt.Printf("Mathematics applicants selected in round %v:\n", i)
				for _, value := range mathematicsApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[2])
				}
				fmt.Println()
			}

			// remove selected Engineering applicants from main applicant list
			applicants = removeApplicants(applicants, mathematicsApplicants)
		}

		// physics
		if len(physicsApplicants) < n {
			newApplicants := filterApplicants(applicants, "Physics", i)

			if debug {
				fmt.Printf("Physics applicants available in round %v:\n", i)
				for _, value := range newApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[0])
				}
				fmt.Println()
			}

			// Sort new Physics applicants by exam (descending order), ties broken by full name
			sort.Slice(newApplicants, func(i, j int) bool {
				meanExamScoreI = (newApplicants[i].ExamScores[2] + newApplicants[i].ExamScores[0]) / 2
				meanExamScoreJ = (newApplicants[j].ExamScores[2] + newApplicants[j].ExamScores[0]) / 2
				deptExamScoreI := newApplicants[i].ExamScores[4]
				deptExamScoreJ := newApplicants[j].ExamScores[4]
				examScoreI := max(meanExamScoreI, deptExamScoreI)
				examScoreJ := max(meanExamScoreJ, deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := newApplicants[i].FirstName + " " + newApplicants[i].LastName
				fullNameJ := newApplicants[j].FirstName + " " + newApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			physicsApplicants = append(physicsApplicants, newApplicants...)

			if len(physicsApplicants) > n {
				physicsApplicants = physicsApplicants[:n]
			}

			// Sort Physics applicants by exam (descending order), ties broken by full name
			sort.Slice(physicsApplicants, func(i, j int) bool {
				meanExamScoreI = (physicsApplicants[i].ExamScores[2] + physicsApplicants[i].ExamScores[0]) / 2
				meanExamScoreJ = (physicsApplicants[j].ExamScores[2] + physicsApplicants[j].ExamScores[0]) / 2
				deptExamScoreI := physicsApplicants[i].ExamScores[4]
				deptExamScoreJ := physicsApplicants[j].ExamScores[4]
				examScoreI := max(meanExamScoreI, deptExamScoreI)
				examScoreJ := max(meanExamScoreJ, deptExamScoreJ)

				if examScoreI != examScoreJ {
					return examScoreI > examScoreJ
				}

				fullNameI := physicsApplicants[i].FirstName + " " + physicsApplicants[i].LastName
				fullNameJ := physicsApplicants[j].FirstName + " " + physicsApplicants[j].LastName
				return fullNameI < fullNameJ
			})

			if debug {
				fmt.Printf("Physics applicants selected in round %v:\n", i)
				for _, value := range physicsApplicants {
					fmt.Print(value.FirstName + " " + value.LastName + " ")
					fmt.Println(value.ExamScores[0])
				}
				fmt.Println()
			}

			// remove selected Engineering applicants from main applicant list
			applicants = removeApplicants(applicants, physicsApplicants)
		}
	}

	outputFile, err := os.Create("biotech.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	writer := bufio.NewWriter(outputFile)

	for _, value := range biotechApplicants {
		meanExamScore := (value.ExamScores[0] + value.ExamScores[1]) / 2
		examScore := max(meanExamScore, value.ExamScores[4])

		_, err1 := fmt.Fprint(writer, value.FirstName+" "+value.LastName+" ")
		if err1 != nil {
			return
		}
		_, err3 := fmt.Fprintf(writer, "%.1f\n", examScore)
		if err3 != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	err = outputFile.Close()
	if err != nil {
		return
	}

	fmt.Println("Biotech")
	for _, value := range biotechApplicants {
		meanExamScore := (value.ExamScores[0] + value.ExamScores[1]) / 2
		examScore := max(meanExamScore, value.ExamScores[4])
		fmt.Print(value.FirstName + " " + value.LastName + " ")
		fmt.Printf("%.1f\n", examScore)
	}

	fmt.Println()

	outputFile, err = os.Create("chemistry.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	writer = bufio.NewWriter(outputFile)

	for _, value := range chemistryApplicants {
		_, err1 := fmt.Fprint(writer, value.FirstName+" "+value.LastName+" ")
		if err1 != nil {
			return
		}
		examScore := max(value.ExamScores[1], value.ExamScores[4])
		_, err3 := fmt.Fprintf(writer, "%.1f\n", examScore)
		if err3 != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	err = outputFile.Close()
	if err != nil {
		return
	}
	fmt.Println("Chemistry")
	for _, value := range chemistryApplicants {
		examScore := max(value.ExamScores[1], value.ExamScores[4])
		fmt.Print(value.FirstName + " " + value.LastName + " ")
		fmt.Printf("%.1f\n", examScore)
	}

	fmt.Println()

	outputFile, err = os.Create("engineering.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	writer = bufio.NewWriter(outputFile)

	for _, value := range engineeringApplicants {
		meanExamScore := (value.ExamScores[2] + value.ExamScores[3]) / 2
		examScore := max(meanExamScore, value.ExamScores[4])
		_, err1 := fmt.Fprint(writer, value.FirstName+" "+value.LastName+" ")
		if err1 != nil {
			return
		}
		_, err3 := fmt.Fprintf(writer, "%.1f\n", examScore)
		if err3 != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	err = outputFile.Close()
	if err != nil {
		return
	}

	fmt.Println("Engineering")
	for _, value := range engineeringApplicants {
		meanExamScore := (value.ExamScores[2] + value.ExamScores[3]) / 2
		examScore := max(meanExamScore, value.ExamScores[4])
		fmt.Print(value.FirstName + " " + value.LastName + " ")
		fmt.Printf("%.1f\n", examScore)
	}

	fmt.Println()

	outputFile, err = os.Create("mathematics.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	writer = bufio.NewWriter(outputFile)

	for _, value := range mathematicsApplicants {
		_, err1 := fmt.Fprint(writer, value.FirstName+" "+value.LastName+" ")
		if err1 != nil {
			return
		}
		examScore := max(value.ExamScores[2], value.ExamScores[4])
		_, err3 := fmt.Fprintf(writer, "%.1f\n", examScore)
		if err3 != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	err = outputFile.Close()
	if err != nil {
		return
	}

	fmt.Println("Mathematics")
	for _, value := range mathematicsApplicants {
		examScore := max(value.ExamScores[2], value.ExamScores[4])
		fmt.Print(value.FirstName + " " + value.LastName + " ")
		fmt.Printf("%.1f\n", examScore)
	}

	fmt.Println()

	outputFile, err = os.Create("physics.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	writer = bufio.NewWriter(outputFile)

	for _, value := range physicsApplicants {
		meanExamScore := (value.ExamScores[0] + value.ExamScores[2]) / 2
		examScore := max(meanExamScore, value.ExamScores[4])

		_, err1 := fmt.Fprint(writer, value.FirstName+" "+value.LastName+" ")
		if err1 != nil {
			return
		}
		_, err3 := fmt.Fprintf(writer, "%.1f\n", examScore)
		if err3 != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	err = outputFile.Close()
	if err != nil {
		return
	}

	fmt.Println("Physics")
	for _, value := range physicsApplicants {
		meanExamScore := (value.ExamScores[0] + value.ExamScores[2]) / 2
		examScore := max(meanExamScore, value.ExamScores[4])
		fmt.Print(value.FirstName + " " + value.LastName + " ")
		fmt.Printf("%.1f\n", examScore)
	}
}

func filterApplicants(applicants []Applicant, priority string, rank int) []Applicant {
	//fmt.Printf("Searching for %v\n", priority)
	var filteredApplicants []Applicant
	for _, applicant := range applicants {
		//fmt.Printf("Applicant priority #%v: %v\n", rank, applicant.Priorities[rank-1])
		if applicant.Priorities[rank-1] == priority {
			filteredApplicants = append(filteredApplicants, applicant)
			//fmt.Printf("Found priority %v\n", applicant.Priorities[rank-1])
		}
	}
	return filteredApplicants
}

func removeApplicants(applicants []Applicant, toRemove []Applicant) []Applicant {
	var updatedApplicants []Applicant
	for _, applicant := range applicants {
		found := false
		for _, remove := range toRemove {
			if applicant.Equals(remove) {
				found = true
				break
			}
		}
		if !found {
			updatedApplicants = append(updatedApplicants, applicant)
		}
	}
	return updatedApplicants
}

func (a Applicant) Equals(other Applicant) bool {
	if a.FirstName != other.FirstName || a.LastName != other.LastName {
		return false
	}
	if len(a.Priorities) != len(other.Priorities) {
		return false
	}
	for i := range a.Priorities {
		if a.Priorities[i] != other.Priorities[i] {
			return false
		}
	}
	return true
}
