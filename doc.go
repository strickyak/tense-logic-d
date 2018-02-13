/*
	An evaluator for a simple Tense (or Temporal) Logic system called D, invented by Robert A. Bull.

	On each day, it either rains or it does not rain.

	The days are Today, Tomorrow, the day after tomorrow, and so forth, into the infinite future.

	The primative Proposition is "R": it is raining today.

	The operators are "T p" for Tomorrow(p), "E p" for Eventually(p), and "F p" for Forevermore(p).

	Tomorrow(p) means p is true tomorrow.

	Eventually(p) means p is true today or tomorrow or some following day (possible more than one day).

	Forevermore(p) means p is true today and tomorrow and all days into the future.


	See http://homepages.mcs.vuw.ac.nz/~rob/papers/modalhist.pdf ( p39: [Robert A.] Bull's Tense Algebra. "[Arthur N.] Prior called this logic D." )
*/
package D
