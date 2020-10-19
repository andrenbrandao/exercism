const colorIndex = color => COLORS.indexOf(color);

export const decodedValue = ([firstColor, secondColor]) =>
  colorIndex(firstColor) * 10 + colorIndex(secondColor);

const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
];
