export const modulus = (n: number, mod: number): number => {
  while (n < 0) return n + mod;
  while (n > mod - 1) return n - mod;
  return n;
};
