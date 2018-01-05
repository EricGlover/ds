package ds

//List is the data type of the Queue 
type List []int

//Queue is a basic queue for ints
//obviously I'll need to extend this later
type Queue interface {
  Enqueue(int) int
  Dequeue() int
  Peek() int
}
//Enqueue adds an item to the Queue, and returns the Q's length
func (q *List) Enqueue(n int) int{
  *q = append(*q, n)
  return len(*q)
}
//Dequeue removes an item from the Q and returns it
//TODO: add an error
func (q *List) Dequeue() int {
  if len(*q) < 1 {
    return 0
  }
  tmp := (*q)[0]
  *q = (*q)[1:]
  return tmp
}
//Peek returns the first item in the Queue
func (q *List) Peek() int {
  return (*q)[0]
}
