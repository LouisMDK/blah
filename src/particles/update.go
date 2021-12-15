package particles

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {

	// Mettre à jour les positions X et Y de chaque particule du système
	for i := 0; i < len(s.Content); i++ {
		s.Content[i].PositionX += s.Content[i].VitesseX
		s.Content[i].PositionY += s.Content[i].VitesseY
	}

	// Met à jour le contenu du système suivanr les valeurs du fichier config
	s.UpdateContent()
}
