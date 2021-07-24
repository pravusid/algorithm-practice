package programmers

type Queue []int

func (q Queue) Add(data int) Queue {
	return append(q, data)
}

func (q Queue) Remove() (Queue, int) {
	return q[1:], q[0]
}

func (q Queue) Peek() int {
	return q[0]
}

func solution(progresses Queue, speeds Queue) []int {
	var result []int

	for len(progresses) != 0 {
		progress(progresses, speeds)

		completed := deployable(&progresses, &speeds)
		if completed > 0 {
			result = append(result, completed)
		}
	}

	return result
}

func progress(progresses Queue, speeds Queue) {
	for i, _ := range progresses {
		progresses[i] += speeds[i]
	}
}

func deployable(progresses *Queue, speeds *Queue) int {
	var completed int
	for _, p := range *progresses {
		if p >= 100 {
			completed++
			pPtr := *progresses
			*progresses, _ = pPtr.Remove()
			sPtr := *speeds
			*speeds, _ = sPtr.Remove()
		} else {
			break
		}
	}
	return completed
}
