export function equals(a1: any[], a2: any[]): boolean {
  if (a1 === a2) {
    return true;
  }
  if ((!a1 && a2) || (a1 && !a2)) {
    return false;
  }

  if (a1.length !== a2.length) {
    return false;
  }

  return a1.every((val, i) => val === a2[i]);
}
