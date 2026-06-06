package testing

import (
	"my_language/ast"
	"reflect"
	"testing"
)

func TestPointers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name: "Simple pointer",
			input: `
                var x = 5
    			var ptr = &x 
				var testResult = *ptr
            `,
			expected: 5,
		},
		{
			name: "Pointer mutation",
			input: `
                var x = 5
    			var ptr = &x 
				*ptr = 7
				var testResult = *ptr
            `,
			expected: 7,
		},
		{
			name: "Pointer on array item",
			input: `
                var numbers = [1, 2, 3]
    			var ptr = &numbers[1]
				var testResult = *ptr
            `,
			expected: 2,
		},
		{
			name: "Pointer on array",
			input: `
                var numbers = [1, 2, 3]
    			var ptr = &numbers[1]
				*ptr = 4
				var testResult = numbers
            `,
			expected: &[]any{1, 4, 3},
		},
		{
			name: "Double Pointer",
			input: `
                var x = 1
				var ptr = &x
				var pptr = &ptr 

				var testResult = **pptr 
            `,
			expected: 1,
		},
		{
			name: "Pointers comparison",
			input: `
                var v1 = 1
				var p1 = &v1
				var p2 = &v1

				var testResult = p1 == p2
            `,
			expected: true,
		},
		{
			name: "Pointers to form linked list",
			input: `

				struct Node {
					val;
					next;
				}

                var node1 = Node { val: 7, next: null }
				var node2 = Node { val: 11, next: null }
				var node3 = Node { val: 10, next: null }
				var node4 = Node { val: 2, next: null }
				var node5 = Node { val: 9, next: null }

				node1.next = node2
				node2.next = node3
				node3.next = node4
				node4.next = node5

				var testResult = node1
            `,
			expected: "Node(7) -> Node(11) -> Node(10) -> Node(2) -> Node(9) -> null",
		},
		{
			name: "Deleting node form linked list",
			input: `

				struct Node {
					val;
					next;
				}


				func deleteSpecificNode(head, nodeToDelete) {
					if (head == nodeToDelete){
						return head.next
					}

					var currentNode = head
					while (currentNode.next != null and currentNode.next != nodeToDelete) {
						currentNode = currentNode.next
					}

					if (currentNode.next == null) {
						return head
					}

					currentNode.next = currentNode.next.next

					return head

				}

                var node1 = Node { val: 7, next: null }
				var node2 = Node { val: 11, next: null }
				var node3 = Node { val: 10, next: null }
				var node4 = Node { val: 2, next: null }
				var node5 = Node { val: 9, next: null }

				node1.next = node2
				node2.next = node3
				node3.next = node4
				node4.next = node5


    			deleteSpecificNode(node1, node2)
				var testResult = node1
            `,
			expected: "Node(7) -> Node(10) -> Node(2) -> Node(9) -> null",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := runInterpreter(tc.input)
			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' missing")
			}

			var actual any = result
			if _, isExpectedString := tc.expected.(string); isExpectedString {
				actual = ast.Stringify(result)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
