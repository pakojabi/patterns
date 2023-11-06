package chainofresposibility

import "fmt"

type Creature struct {
	Name string
	Attack, Defense int
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s: %d/%d", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

func (cm *CreatureModifier) Add(m Modifier) {
	if cm.next != nil {
		cm.next.Add(m)
	} else {
		cm.next = m
	}
}

func (cm *CreatureModifier) Handle() {
	if cm.next != nil {
		cm.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(creature *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: creature}}
}

func (dam *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling creature's attack")
	dam.creature.Attack *= 2
	dam.CreatureModifier.Handle()
}

type DoubleDefenseModifier struct {
	CreatureModifier
}

func NewDoubleDefenseModifier(creature *Creature) *DoubleDefenseModifier {
	return &DoubleDefenseModifier{CreatureModifier{creature: creature}}
}

func (ddm *DoubleDefenseModifier) Handle() {
	fmt.Println("Doubling creature's defense")
	ddm.creature.Defense *= 2
	ddm.CreatureModifier.Handle()
}

func RunMethodChain() {
	goblin := NewCreature("Goblin", 1, 1)
	modifiers := NewCreatureModifier(goblin)
	modifiers.Add(NewDoubleAttackModifier(goblin))
	modifiers.Add(NewDoubleDefenseModifier(goblin))
	modifiers.Add(NewDoubleDefenseModifier(goblin))

	modifiers.Handle()

	fmt.Println(goblin.String())
}





