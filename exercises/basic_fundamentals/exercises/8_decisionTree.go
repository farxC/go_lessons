package exercises

type Node struct {
	question string
	yes      *Node
	no       *Node
	action   string
}

type DecisionTree struct {
	root *Node
}

func BuildDecisionTree() *DecisionTree {
	root := &Node{}

	root.question = "Is it raining?"

	root.yes = &Node{question: "Is it a weekday?"}

	root.no = &Node{question: "Do you want to socialize?"}

	root.no.yes = &Node{action: "Visit friends."}

	root.no.no = &Node{action: "Watch a movie."}

	root.yes.yes = &Node{
		action: "Go to work.",
	}

	root.yes.no = &Node{question: "Is it a holiday?"}

	root.yes.no.yes = &Node{action: "Read a book."}

	root.yes.no.no = &Node{action: "Stay inside."}

	return &DecisionTree{root: root}

}

func (dt *DecisionTree) Traverse(answers []bool) string {

	tree := BuildDecisionTree()
	current := tree.root

	for i := 0; i < len(answers); i++ {
		if current == nil {
			break
		}
		if answers[i] {
			current = current.yes
		} else {
			current = current.no
		}
	}

	if current != nil {
		return current.action
	}
	return "No decision could be made."
}
