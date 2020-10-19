const dnaNucleotideComplement = {
  G: 'C',
  C: 'G',
  T: 'A',
  A: 'U',
};

const dnaToRnaNucleotide = nucleotide => {
  return dnaNucleotideComplement[nucleotide];
};

export const toRna = dnaStrand => {
  if (!dnaStrand) return '';

  return dnaToRnaNucleotide(dnaStrand[0]) + toRna(dnaStrand.substr(1));
};
