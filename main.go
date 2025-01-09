package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
	"glob/mod" // Assurez-vous que ce package est défini
)

type Stat struct {
	Name   string
	PV     int
	MaxPV  int
	Force  int
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialise l'aléatoire

	fmt.Println("--------------------------------")
	fmt.Println("Bienvenue dans le jeu !")
	fmt.Println("--------------------------------")
	time.Sleep(1 * time.Second)
	fmt.Println("--------------------------------")
	fmt.Println("Voulez-vous commencer une nouvelle partie?")
	fmt.Println("[Y]/[N]")
	fmt.Println("--------------------------------")

	var debut string
	fmt.Scanln(&debut)
	if debut == "Y" || debut == "y" {
		fmt.Println("--------------------------------")
		fmt.Println("Début de la partie !")
		fmt.Println("Statut de départ : (60 PV) (35 Soin) (15 Force) ")
		fmt.Println("--------------------------------")
		jeu() // appel de la fonction jeu
	} else {
		fmt.Println("--------------------------------")
		log.Fatalf("Ok, à bientôt !")
	}
}

func attaque(att, def *Stat) {
	time.Sleep(1 * time.Second)
	fmt.Println("--------------------------------")
	fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", att.Name, att.Force, def.Name)
	fmt.Println("--------------------------------")
	def.PV -= att.Force
	if def.PV < 0 {
		def.PV = 0
	}
}

func soin(state *Stat) {
	time.Sleep(1 * time.Second)
	soin1 := int(math.Min(float64(30), float64(state.MaxPV-state.PV)))
	state.PV += soin1
	fmt.Println("--------------------------------")
	fmt.Printf("%s se soigne de %d PV. Ses PV actuels sont : %d.\n", state.Name, soin1, state.PV)
	fmt.Println("--------------------------------")
}

func estEnVie(state *Stat) bool {
	return state.PV > 0
}

func jeu() {
	link := Stat{Name: "Link", PV: 60, MaxPV: 60, Force: 15}

	for etage := 1; etage <= 10; etage++ {
		var mob Stat
		var boss Stat

		if etage < 10 { // mob jusque à boss
			randommob := rand.Intn(11) + 1 // mob aléatoire
			mobData, err := mod.Arrivagedemob(strconv.Itoa(randommob)) // conversion string
			if err != nil {
				log.Fatalf("Erreur lors de la génération du monstre : %v", err)
			}
			mob = convstatmob(mobData) // monstre converti

			fmt.Println("--------------------------------")
			fmt.Printf("Étage %d : Vous affrontez un %s (%d PV, %d Force).\n", etage, mob.Name, mob.PV, mob.Force)
			fmt.Println("--------------------------------")

		} else { 
			randomBoss := rand.Intn(3) + 1 // boss aléatoire
			bossData, err := mod.Arrivagedeboss(strconv.Itoa(randomBoss)) // conversion string
			if err != nil {
				log.Fatalf("Erreur lors de la génération du boss : %v", err)
			}
			boss = convstatboss(bossData) // conversion boss en stat

			fmt.Println("--------------------------------")
			fmt.Printf("Étage %d : Vous affrontez le boss final : %s (%d PV, %d Force).\n", etage, boss.Name, boss.PV, boss.Force)
			fmt.Println("--------------------------------")
		}

		for estEnVie(&link) && estEnVie(&mob) { // combat
			var action int
			time.Sleep(1 * time.Second)
			fmt.Println("--------------------------------")
			fmt.Println("Que voulez-vous faire ?")
			fmt.Println("1. Attaquer")
			fmt.Println("2. Se soigner")
			fmt.Println("--------------------------------")
			fmt.Scan(&action)

			switch action {
			case 1:
				attaque(&link, &mob)
			case 2:
				soin(&link)
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("--------------------------------")
				fmt.Println("Choix invalide, veuillez choisir 1 ou 2.")
				fmt.Println("--------------------------------")
				continue
			}

			if estEnVie(&mob) {
				attaque(&mob, &link)
			}
			time.Sleep(1 * time.Second)
			fmt.Println("--------------------------------")
			fmt.Printf("PV de %s : %d/%d\n", link.Name, link.PV, link.MaxPV)
			fmt.Printf("PV de %s : %d/%d\n", mob.Name, mob.PV, mob.MaxPV)
			fmt.Println("--------------------------------")
		}

		if !estEnVie(&link) {
			time.Sleep(1 * time.Second)
			fmt.Println("--------------------------------")
			fmt.Println("Vous avez perdu.")
			fmt.Println("--------------------------------")
			return
		}

		if !estEnVie(&mob) {
			time.Sleep(1 * time.Second)
			fmt.Println("--------------------------------")
			fmt.Printf("%s est mort !\n", mob.Name)
			fmt.Println("--------------------------------")
		}

		if etage == 10 && !estEnVie(&mob) { // fin partie
			time.Sleep(1 * time.Second)
			fmt.Println("--------------------------------")
			fmt.Println("Félicitations, vous avez vaincu le boss !")
			fmt.Println("--------------------------------")
			return
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("--------------------------------")
	fmt.Println("Félicitations, le monde se portera mieux maintenant !")
	fmt.Println("--------------------------------")
}

func convstatmob(mob mod.Mob) Stat { // mob converti en stat
	return Stat{
		Name:   mob.Name,
		PV:     mob.PV,
		MaxPV:  mob.MaxPV,
		Force:  mob.Force,
	}
}

func convstatboss(boss mod.Boss) Stat { // boss converti en stat
	return Stat{
		Name:   boss.Name,
		PV:     boss.PV,
		MaxPV:  boss.MaxPV,
		Force:  boss.Force,
	}
}
