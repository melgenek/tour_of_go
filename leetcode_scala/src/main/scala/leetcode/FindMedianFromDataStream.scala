package leetcode

object FindMedianFromDataStream extends App {

  val finder = new MedianFinder

  finder.addNum(6)
  println(finder.findMedian() == 6)

  finder.addNum(10)
  println(finder.findMedian() == 8)

  finder.addNum(2)
  println(finder.findMedian() == 6)

  finder.addNum(6)
  println(finder.findMedian() == 6)

  finder.addNum(5)
  println(finder.findMedian() == 6)

  finder.addNum(0)
  println(finder.findMedian() == 5.5)

  finder.addNum(6)
  println(finder.findMedian() == 6)

  finder.addNum(3)
  println(finder.findMedian() == 5.5)

  finder.addNum(1)
  println(finder.findMedian() == 5)

  finder.addNum(0)
  println(finder.findMedian() == 4)

  finder.addNum(0)
  println(finder.findMedian() == 3)


}

class MedianFinder {

  import scala.collection.mutable

  private val low = mutable.PriorityQueue.empty[Int]
  private val high = mutable.PriorityQueue.empty[Int](Ordering[Int].reverse)

  def addNum(num: Int) {
    if (low.isEmpty || low.head >= num) {
      low.addOne(num)
    } else {
      high.addOne(num)
    }

    if (low.length - high.length >= 2) {
      high.addOne(low.dequeue())
    } else if (high.length - low.length >= 2) {
      low.addOne(high.dequeue())
    }
  }

  def findMedian(): Double = {
    if (low.length > high.length) {
      low.head.toDouble
    } else if (low.length < high.length) {
      high.head.toDouble
    } else {
      (low.head + high.head).toDouble / 2
    }
  }

}

