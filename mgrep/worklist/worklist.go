package worklist

type Entry struct {
	Path string
}

type Worklist struct {
	job chan Entry
}

func (job *Worklist) Add(work Entry) {
	job.job <- work
}

func (job *Worklist) Next() Entry {
	work := <-job.job
	return work
}

func New(buffSize int) Worklist {
	return Worklist{
		job: make(chan Entry, buffSize),
	}
}

func NewJob(path string) Entry {
	return Entry{path}
}

func (workList *Worklist) Finalize(workerNumbers int) {
	for i := 0; i < 10; i++ {
		workList.Add(Entry{Path: ""})
	}
}
