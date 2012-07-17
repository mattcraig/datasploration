;; ( 6 (3 (2) (4)) (9 (7) (8)))

(defun tcreate (L)
  "from this list create a tree."
  (if (null L)
	nil
	(tadd (first (last L)) (tcreate (reverse (rest (reverse L)))))
	)
  )

(defun tadd (v N)
  " add value v to the tree starting at node N. add as leaf node"
  (cond
	((null N)   (list v nil nil))
	((>= v (first N))   (tbalance (list (first N) (tleft N) (tadd v (tright N)))))
	( t   (tbalance (list (first N) (tadd v (tleft N)) (tright N) )))
	)
  )

(defun tadd-top (v N)
  " add value v to the tree starting at node N.  Add as the top node"
  (cond
	((null N)   (list v nil nil))
	((>= v (first N))   (list v N nil))
	( t   (list v nil N))
	)
  )

(defun tdepth (N)
  "the depth of the deepest part of the tree"
  (if (null N)
	0
	(1+ (max (tdepth (tleft N)) (tdepth (tright N))))
	)
  )

(defun tbalance (N)
  "from whichever side is greater depth shift up the node closest to the top node"
  (cond
	((null N)   nil)
	((and (null (tleft N)) (> (tdepth (tright N)) 1))
	 (tadd (first N) (tbalance (tright N)))
	 )
	((and (null (tright N)) (> (tdepth (tleft N)) 1))
	 (tadd (first N) (tbalance (tleft N)))
	 )
	((> (- (tdepth (tleft N)) (tdepth (tright N))) 1)
	 ;; left side is already too deep
	 ;; find the next closest < value and move it up to top with this node as its right
	 (list
	   (tmax (tleft N))
	   (tbalance (tremoveleaf (tmax (tleft N)) (tleft N)))
	   (tbalance (list (first N) nil (tright N)))
	   )
	 )
	((> (- (tdepth (tright N)) (tdepth (tleft N))) 1)
	 ;; right side is already too deep
	 ;; find the next closest > value and move it up to top with this node as its left
	 (list
	   (tmin (tright N))
	   (tbalance (list (first N) (tleft N) nil))
	   (tbalance (tremoveleaf (tmin (tright N)) (tright N)))
	   )
	 )
	(t   N)
	)
  )

(defun tremoveleaf (v N)
  "find v in this tree and remove it"
  (cond
	((null N)   nil)
	((= v (first N))
	 (if (and (null (tleft N)) (null (tright N)))
	   nil
	   N
	   )
	 )
	(t    (list (first N) (tremoveleaf v (tleft N)) (tremoveleaf v (tright N))))
	)

  )

(defun tmax (N)
  "get the maximum value in the tree starting at this node"
  (cond
	((null N)   nil)
	((null (tright N))    (first N))
	(t   (tmax (tright N)))
	)
  )

(defun tmin (N)
  "get the minimum value in the tree starting at this node"
  (cond
	((null N)   nil)
	((null (tleft N))    (first N))
	(t   (tmin (tleft N)))
	)
  )


(defun tleft (N)
  "fetch the left child of tree node N"
  (if (null N)
	nil
	(nth 1 N)
	)
  )

(defun tright (N)
  "fetch the right child of tree node N"
  (if (null N)
	nil
	(nth 2 N)
	)
  )

(defun tpreorder (N)
  "return a list ordered as in a breadth first search"
  (cond
	((and (null (tleft N)) (null (tright N)))	(list (first N)))
	((null (tleft N))
	 (cons (first N) (tpreorder (tright N)))
	 )
	((null (tright N))
	 (cons (first N) (tpreorder (tleft N)))
	 )
	(t  (cons (first N) (append (tpreorder (tleft N)) (tpreorder (tright N)))))
	)
  )

(defun tpostorder (N)
  "return a list ordered as in a depth first search"
  (cond
	((and (null (tleft N)) (null (tright N)))	(list (first N)))
	((null (tleft N))
	 (append (tpostorder (tright N)) (list (first N)))
	 )
	((null (tright N))
	 (append (tpostorder (tleft N)) (list (first N)))
	 )
	(t    (append (append (tpostorder (tleft N)) (tpostorder (tright N))) (list  (first N))) )
	)
  )

(setq tree '( 6 (3 (2) (4)) (9 (7) (8))))
(tleft tree)