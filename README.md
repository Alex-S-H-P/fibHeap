# FibHeap

A fibonacci heap.

Data is stored in a forest of tree, that is regularly cleaned to minimize the ammount of trees. 
Every bit of data stored has its own priority, and the heap is optimized to pull from the **highest** priority first.

# Heap and CHeap

In most use-cases, it is recommanded to use CHeap and not Heap, as CHeap has the ability to pull nodes from their content.

* `Heap[P, T]` can be applied for any Number (except complex) `P` and any type `T` of data to be stored.
* `CHeap[P, T]` can be applied for any Number (except complex) `P` and any type `T` of data **that allows equality** to be stored.

Both implement these simple functions : 
* `GetMax` which returns the node with the highest priority
* `ExtractMax` which disolves the node with the highest priority, deleting it from memory (For example because it's its turn to be handled)
* `Insert` and InsertNodes, which add nodes
* `IncreasePriority`, which affects the priority of a node in particular
* `Merge` that emtpies one Heap into another
