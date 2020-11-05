const WEEKDAY = [
  'Sunday',
  'Monday',
  'Tuesday',
  'Wednesday',
  'Thursday',
  'Friday',
  'Saturday',
];

const WEEK = {
  first: 1,
  second: 8,
  third: 15,
  fourth: 22,
  teenth: 13,
  last: -6,
};

export const meetup = (year, month, modifier, weekday) => {
  const firstDay = WEEK[modifier];
  const lastDay = firstDay + 7;
  const currentMonth = firstDay < 0 ? month : month - 1;

  for (let day = firstDay; day < lastDay; day++) {
    const dateFound = new Date(year, currentMonth, day);
    if (dateFound.getDay() === WEEKDAY.indexOf(weekday)) {
      return dateFound;
    }
  }
  throw new Error('Could not find date.');
};
