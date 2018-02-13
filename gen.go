package D

func Generate(n uint) []Story {
	var stories []Story
	var i uint
	for i = 0; i < (1 << n); i++ {
		v := make(Story, n)
		var p uint
		for p = 0; p < n; p++ {
			if ((i >> p) & 1) == 1 {
				v[n-p-1] = true
			} else {
				v[n-p-1] = false
			}
		}
		stories = append(stories, v)
	}
	return stories
}
