;;estcroissante fais chier la cadrablevalue...avec >. J'ai toujours des blèmes avec cet opérateur...(cf minliste...)
(define (estcroissante? L)
  (if (<= (length L) 1)
      #t
      (if (< (car L) (cadr L))
          (estcroissante? (cdr L))
          #f)))
  
;;cette fonction fait ce qui est demandé pour nieme bis...
(define (nieme L x)
  (if (pair? L)
      (if (= x 0)
          (car L)
          (nieme (cdr L) (- x 1)))
      #f))

;;les 3 dernières questions...

     
  

