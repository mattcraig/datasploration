class DSSet :

    def __init__(self, items = []):
        self.set_data = {}

        for item in items :
            self.add(item)

    def add(self, item) :
        set_item = self._get_key(item)
        self.set_data[set_item] = item

    def remove(self, item) :
        set_item = self._get_key(item)
        if set_item in self.set_data :
            del self.set_data[set_item]

    def cardinality(self) :
        return len(self.set_data)

    def union(self, otherset) :
        ret_set = self.set_data.copy()
        for item in otherset :
            set_item = self._get_key(item)
            if not set_item in ret_set:
                ret_set[set_item] = set_item

        return DSSet(ret_set)

    def intersection(self, otherset) :
        ret_set = []
        for item in otherset :
            set_item = self._get_key(item)
            if set_item in self.set_data :
                ret_set.append(set_item)

        return DSSet(ret_set)

    def set_theoretic_difference(self, otherset) :
        ret_set = self.set_data.copy()
        for item in otherset :
            set_item = self._get_key(item)
            if set_item in self.set_data :
               del ret_set[set_item]

        return DSSet(ret_set)

    def symmetric_difference(self, otherset) :
        ret_set = self.set_data.copy()
        for item in otherset :
            set_item = self._get_key(item)
            if set_item in self.set_data :
                del ret_set[set_item]
            else :
                ret_set[set_item] = item

        return DSSet(ret_set)

    def is_member(self, item) :
        set_item = self._get_key(item)
        if set_item in self.set_data :
            return True

        return False

    def is_subset_of(self, otherset) :
        for item in self.set_data :
            if not item in otherset :
                return False

        return True

    def is_superset_of(self, otherset) :
        for item in otherset :
            if not item in self.set_data :
                return False

        return True


    def __iter__(self) :
        return self.set_data.itervalues()

    def __str__(self) :
        return str(self.set_data.values())

    def _get_key(self, item) :
        return str(item)

    def i_do_not_know_how_to_copy_a_dict(self):
        ret_dict = {}
        for k,v in self.set_data.iteritems() :
            ret_dict[k] = v

        return ret_dict

sniggle = DSSet(['hobo', 9 , 18.9, "18.9"])
biggle = DSSet(['goat', 9 , ['island', 9]])

print sniggle
print biggle
nigh = sniggle.intersection(biggle)
print "intersection"
print nigh
print nigh.cardinality()
nigh = sniggle.union(biggle)
print "union"
print nigh
print nigh.cardinality()

nigh = sniggle.set_theoretic_difference(biggle)
print "set_theor"
print nigh
nigh = biggle.symmetric_difference(sniggle)
print "symm"
print nigh

print sniggle
print biggle