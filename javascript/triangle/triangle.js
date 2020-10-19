export class Triangle {
  constructor(sideOne, sideTwo, sideThree) {
    this.sideOne = sideOne;
    this.sideTwo = sideTwo;
    this.sideThree = sideThree;
  }

  hasPositiveSides() {
    return this.sideOne > 0 && this.sideTwo > 0 && this.sideThree > 0;
  }

  validatesTriangleInequality() {
    if (this.sideOne + this.sideTwo < this.sideThree) return false;
    if (this.sideOne + this.sideThree < this.sideTwo) return false;
    if (this.sideTwo + this.sideThree < this.sideOne) return false;
    return true;
  }

  isTriangle() {
    if (this.hasPositiveSides() && this.validatesTriangleInequality()) {
      return true;
    }
    return false;
  }

  isEquilateral() {
    if (!this.isTriangle()) return false;

    return this.sideOne === this.sideTwo
      ? this.sideOne === this.sideThree
      : false;
  }

  isIsosceles() {
    if (!this.isTriangle()) return false;

    return (
      this.sideOne === this.sideTwo ||
      this.sideOne === this.sideThree ||
      this.sideTwo === this.sideThree
    );
  }

  isScalene() {
    if (!this.isTriangle()) return false;

    return this.sideOne !== this.sideTwo
      ? this.sideOne !== this.sideThree
      : false;
  }
}
