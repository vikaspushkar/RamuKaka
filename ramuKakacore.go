package ramukaka

import (
	"fmt"
	"strings"
)

//Ramukaka is a variable of NeuNet.
var Ramukaka NeuNet

//Init initalizes
func Init() {
	Ramukaka.endhere = false
}

//Teach will add the element
func Teach(QorA string) (oops bool, err error) {
	oops, err = add(QorA)
	return false, nil
}
func processInput(QorA string) string {
	a := strings.ToLower(strings.TrimSpace(QorA))
	b := strings.Replace(a, " ", "", -1)
	return b
}

//Show dispaly the tree
func Show(head *NeuNet) {
	//fmt.Print(Ramukaka)
	//head := &Ramukaka
	//fmt.Print("#\n")

	for ii := 0; ii < 36; ii++ {
		if head != nil && head.value[ii] != 0 {
			printnode(head, ii)
			Show(head.next[ii])
		}
	}
	fmt.Print("\n")
}
func printnode(head *NeuNet, index int) {
	fmt.Printf("%c", head.value[index])
	/*if head.endhere {
		fmt.Printf("EndHere is true\n")
	} else {
		fmt.Printf("EndHere is false\n")
	}
	for a, b := range head.qNa {
		fmt.Printf("Answere for %s is %v \n", a, b)
	}
	fmt.Printf("Referral =%d \n", head.frequentreferral[index])
	fmt.Printf("Next =%v\n", head.next[index])
	fmt.Printf("Value=%c\n", head.value[index])*/
}
func add(QorA string) (oops bool, err error) {
	qNa := processInput(QorA)
	var whereAt = 0
	head := &Ramukaka
	hash := hash(qNa)
	whereAt = findIndex(qNa[0])
	fmt.Printf(" Adding at %d  %c\n", whereAt, 97+whereAt)
	if head.value[whereAt] == 0 {
		head.value[whereAt] = qNa[0]
		if head.qNa == nil {
			head.qNa = make(map[string]*NeuNet)
			head.qNa[hash] = nil // need to place answer here.
		} else {
			head.qNa[hash] = nil
		}
		head.frequentreferral[whereAt] = 1
	} else {
		head.frequentreferral[whereAt]++
	}
	if len(qNa) == 1 {
		head.endhere = true
	} else {
		head.endhere = false
	}
	qNa = qNa[1:]
	for ii := 0; ii < len(qNa); ii++ {
		prevwhereAt := whereAt
		if head.next[prevwhereAt] == nil {
			head.next[prevwhereAt] = &NeuNet{} // new node
		}
		whereAt = findIndex(qNa[ii])
		if head.next[prevwhereAt].value[whereAt] == 0 {
			head.next[prevwhereAt].value[whereAt] = qNa[ii]
			if head.next[prevwhereAt].qNa == nil {
				head.next[prevwhereAt].qNa = make(map[string]*NeuNet)
				head.next[prevwhereAt].qNa[hash] = nil // need to place answer here.
			} else {
				head.next[prevwhereAt].qNa[hash] = nil
			}
			head.next[prevwhereAt].frequentreferral[whereAt] = 1
		} else {
			head.next[prevwhereAt].frequentreferral[whereAt]++
		}
		if len(qNa) == ii+1 {
			head.next[prevwhereAt].endhere = true
		} else {
			head.next[prevwhereAt].endhere = false
		}
		head = head.next[prevwhereAt]
	}
	return false, nil
}

//Forget deletes the statements
func Forget(statement string) (oops bool, err error) {
	return delete(statement)
}
func delete(QorA string) (oops bool, err error) {
	head := &Ramukaka
	qNa := processInput(QorA)
	if len(qNa) <= 0 {
		fmt.Printf("Nothing to delete")
		return true, nil
	}
	var whereAt int
	output, _ := find(qNa)
	if output {
		fmt.Printf("the %s %s", qNa, " Does not exists to delete")
		return true, nil
	}
	fmt.Printf("Lets delete the statement %s\n", qNa)
	for jj := 0; jj < len(qNa); jj++ {
		//total()
		whereAt = findIndex(qNa[jj])
		if head.frequentreferral[whereAt] == 0 {
			fmt.Printf("Statement letter %c %s\n", whereAt+97, " Does not exists")
			head.next[whereAt] = nil
			break
		} else {
			fmt.Printf("Reducing statement letter %c \n", whereAt+97)
			head.frequentreferral[whereAt]--
			if head.frequentreferral[whereAt] == 0 {
				fmt.Printf("Deleting statement letter %c \n", whereAt+97)
				temp := head.next[whereAt]
				head.next[whereAt] = nil
				head.value[whereAt] = 0
				head.endhere = false
				head = temp

			} else {
				fmt.Printf("moving from statement letter %c \n", whereAt+97)

				head = head.next[whereAt]
			}
		}

	}
	return false, nil
}

//Ask does find
func Ask(QorA string) (oops bool, err error) {
	return find(QorA)
}
func find(QorA string) (oops bool, err error) {
	var whereAt = 0
	head := &Ramukaka
	qNa := processInput(QorA)
	for ii := 0; ii < len(qNa); ii++ {
		whereAt = findIndex(qNa[ii])
		fmt.Printf("searching %c \n", qNa[ii])
		if head.value[whereAt] != qNa[ii] || head == nil {
			return true, nil
		}
		head = head.next[whereAt]
	}
	return false, nil
}

//ChangeMyMind edits the statements
func ChangeMyMind(OriginQorA, NewQorA string) (oops bool, err error) {
	return edit(OriginQorA, NewQorA)
}

func edit(OriginQorA, NewQorA string) (oops bool, err error) {
	head := &Ramukaka
	var prevTemp *NeuNet
	qNa := processInput(OriginQorA)
	qNaNew := processInput(NewQorA)
	qNalen := len(qNa)
	qNaNewlen := len(qNaNew)
	var ii, jj, samelen, whereAt int
	if qNalen > qNaNewlen {
		fmt.Printf("Original is longer than New statement")
		samelen = qNaNewlen
	} else {
		fmt.Printf("New is longer than Original statement")
		samelen = qNalen
	}
	prevTemp = head
	//total()
	for ii = 0; ii < samelen; ii++ {
		whereAt = findIndex(qNa[ii])
		if head.value[whereAt] == qNaNew[ii] {
			fmt.Printf("same char %c to %c \n", qNa[ii], qNaNew[ii])
			prevTemp = head
			head = head.next[whereAt]
		} else {
			break
		}
	}
	prevTemp = head
	fmt.Printf("Deleting Original statement\n")
	for jj = ii; jj < qNalen; jj++ {
		//total()
		whereAt = findIndex(qNa[jj])
		if head.frequentreferral[whereAt] == 0 {
			fmt.Printf("Original statement letter %c %s\n", whereAt+97, " Does not exists")
			head.next[whereAt] = nil
			break
		} else {
			fmt.Printf("Reducing Original statement letter %c \n", whereAt+97)
			head.frequentreferral[whereAt]--
			if head.frequentreferral[whereAt] == 0 {
				fmt.Printf("Deleting Original statement letter %c \n", whereAt+97)
				temp := head.next[whereAt]
				head.next[whereAt] = nil
				head.value[whereAt] = 0
				head.endhere = false
				head = temp

			} else {
				fmt.Printf("moving from Original statement letter %c \n", whereAt+97)

				head = head.next[whereAt]
			}
		}

	}
	fmt.Printf("Adding New statement  \n")

	for jj = ii; jj < qNaNewlen; jj++ {
		//total()
		whereAt = findIndex(qNaNew[jj])
		fmt.Printf("Increasing new statement letter %c \n", whereAt+97)
		if prevTemp.value[whereAt] == 0 && prevTemp.next[whereAt] == nil {
			prevTemp.next[whereAt] = &NeuNet{} // new node
			fmt.Printf("Letter %c %s\n", whereAt+97, "Does not exists")
			prevTemp.frequentreferral[whereAt] = 1
			prevTemp.value[whereAt] = qNaNew[jj]
			if prevTemp.qNa == nil {
				prevTemp.qNa = make(map[string]*NeuNet)
				hash := hash(qNaNew)
				prevTemp.qNa[hash] = nil // need to place answer here.
			} else {
				hash := hash(qNaNew)
				prevTemp.qNa[hash] = nil // need to place answer here.
			}
		} else {
			prevTemp.frequentreferral[whereAt]++
		}
		prevTemp = prevTemp.next[whereAt]
	}
	total()
	return false, nil
}
func total() {
	fmt.Printf("total()  \n")
}

func deleteANodeChar(chartodel int, node *NeuNet) (oops bool, err error) {
	fmt.Printf("deleting %c\n", chartodel+97)
	if node == nil || node.value[chartodel] == 0 {
		fmt.Printf("deleting node is nil\n")
		return true, nil
	}
	node.frequentreferral[chartodel]--
	if node.frequentreferral[chartodel] == 0 {

		node.next[chartodel] = nil
	}
	return false, nil
}
func addANodeChar(chartoadd int, node *NeuNet) (oops bool, err error) {
	fmt.Printf(" Adding node %c\n", chartoadd+97)
	if node == nil {
		fmt.Printf(" Adding node is nil\n")
		return true, nil
	}
	if node.frequentreferral[chartoadd] == 0 {

	}
	node.value[chartoadd] = byte(chartoadd)

	return false, nil
}
