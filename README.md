fmt.PrintF format:
  %v = est le formatteur pas defaut, est polymorphe et affiche la variable dans
  son format par defaut

  %q = est le specifier plaçant des quotes autour d'une chaine de charactère
  %T = specifier affichant le type de la variable
  %x.xf = specifier affichant un float (nombre de charactere symbolisé par x)

fmt.Sprintf: 
  enregistre des chaine de caractere formatté dans une varibale

Array: 
  var ages2 = [3]int{20, 22, 32} = tout comme en C la taille est fixe

Slice:
  Fonctionne sur la base d'un array mais permet de modifier le taille 
  et met a disposition la fonction append(slice, element)
  qui renvoie un tableau contenant le nouvelle element
  var scores = []int{1, 2, 3}
  scores2 := append(scores, 100)

Range: 
  Permet de faire une recopie de tableau a un indice donné borne sup non incluse
  rangeOne := names[1:3]
	rangeTwo := names[0:]
	rangeThree := names[:2]
  NOn destrucive et retourn un slice, ainsi on peut creer un slice a
  a partir d'un array avec un range