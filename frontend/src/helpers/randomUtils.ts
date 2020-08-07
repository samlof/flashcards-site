/**
 * Shuffles array in place.
 * @param {Array} a items An array containing the items.
 */
export function shuffle<T>(a: T[]): T[] {
  let j, x, i;
  for (i = a.length - 1; i > 0; i--) {
    j = Math.floor(Math.random() * (i + 1));
    x = a[i];
    a[i] = a[j];
    a[j] = x;
  }
  return a;
}

/**
 * Gets a random integer between start and end. End is not included
 * @param start Start of range
 * @param end End of range
 */
export function randInt(start: number, end: number): number {
  const length = end - start;
  return Math.floor(Math.random() * length + start);
}
