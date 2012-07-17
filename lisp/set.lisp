(defun sadd (i S)
  "add item i to set S"
  (cond
	((null S) (cons i nil))
	((member i S) S)
	(t (cons i S))
	)
  )

(defun screate (L)
  "create a set by taking unique elements of list L"
  (if (null L)
	nil
	(sadd (first L) (screate (rest L)))
	)
  )

(defun sremove (i S)
  "remove item i from set S"
  (cond
	((null S) nil)
	((member i S) 
	 (cond
	   ((eq i (first S)) (rest S))
	   (t (cons (first S) (sremove i (rest S))))
	 ))
	(t S)
	)
  )

(defun scardinality (S)
  "size of the set"
  (if (null S)
	0
	(list-length S))
)

(defun sunion (S1 S2) 
  "union of S1 and S2"
  (cond
	((and (null S1) (null S2)) nil)
	((null S1) S2)
	((null S2) S1)
	(t (sadd (first S1) (sunion (rest S1) S2)))
	)
  )

(defun sintersection (S1 S2)
  "the intersection of sets S1 and S2"
  (cond
	((or (null S1) (null S2))  nil)
	;; ((member (first S1) S2)  (sadd (first S1) (sintersection (sremove (first S1) S2) (rest S1))))
	((member (first S1) S2)  (sadd (first S1) (sintersection (rest S1) S2)))
	(t (sintersection (rest S1) S2))
	)
  )

(defun set-theoretic-diff (S1 S2)
  "everything from S1 not also in S2"
  (cond
	((null S1)   nil)
	((null S2)   S1)
	((member (first S1) S2)  (set-theoretic-diff (rest S1) S2))
	(t  (sadd (first S1)   (set-theoretic-diff (rest S1) S2)))
	)
  )

(defun symmetric-difference (S1 S2)
  "all items in both sets, not in the intersection of both sets"
  (cond
	((null S1)   S2)
	((null S2)   S1)
	(t (set-theoretic-diff (sunion S1 S2) (sintersection S1 S2)))
	)
  )

(defun ssubsetofp (S1 S2)
  " is S1 a subset of S2"
  (cond
	;; we're going to say that the empty set is a subset of the empty set
	((null S1)  t)
	((member (first S1) S2)   (ssubsetofp (rest S1) S2))
	(t   nil)
	)
  )