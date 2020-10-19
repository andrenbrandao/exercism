export class Clock {
  constructor(hours, minutes = 0) {
    this.hours = 0;
    this.minutes = 0;
    this.setTime(hours, minutes);
  }

  toString() {
    return `${this.formatTime(this.hours)}:${this.formatTime(this.minutes)}`;
  }

  plus(minutes) {
    this.addMinutes(minutes);
    return this;
  }

  minus(minutes) {
    this.addMinutes(-minutes);
    return this;
  }

  equals(clock) {
    return this.hours === clock.hours && this.minutes === clock.minutes;
  }

  formatTime(time) {
    return time > 9 ? time : '0' + time;
  }

  setTime(hours, minutes) {
    this.addHours(hours);
    this.addMinutes(minutes);
  }

  addHours(hours) {
    const partialHours = (this.hours + hours) % 24;
    this.hours = partialHours < 0 ? partialHours + 24 : partialHours;
  }

  addMinutes(minutes) {
    const partialMinutes = (this.minutes + minutes) % 60;
    const hoursRemaining = parseInt((this.minutes + minutes) / 60);

    if (partialMinutes < 0) {
      this.addHours(-1 + hoursRemaining);
      this.minutes = partialMinutes + 60;
    } else {
      this.addHours(hoursRemaining);
      this.minutes = partialMinutes;
    }
  }
}
