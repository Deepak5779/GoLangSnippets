//go:generate stringer -type=Pillg
package painkiller

type Pill int

const (
	Placebo Pill = 7
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
