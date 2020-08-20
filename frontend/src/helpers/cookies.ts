export const cookieValue = (name: string): string | undefined => {
  const cookie = document.cookie
    .split("; ")
    .find((row) => row.startsWith(name));
  if (!cookie) return undefined;

  return cookie.split("=")[1];
};
