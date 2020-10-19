export const reverseString = (text) => {
  let reversedText = '';

  for (let i = text.length - 1; i >= 0; i--) {
    reversedText += text[i];
  }

  return reversedText;
};
