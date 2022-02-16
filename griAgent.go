package griagent

type GriAgent struct {
	arch procArchType
}

func New() *GriAgent {
	cpu, _ := getProcessorInfo()
	return &GriAgent{
		arch: cpu,
	}

}
