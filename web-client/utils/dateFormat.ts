function assertDate(date): Date {
  return date instanceof Date ? date : new Date(date);
}

export function formatDate(date): string {
  date = assertDate(date);
  const d = date.getDate();
  const m = date.getMonth() + 1;
  const y = date.getFullYear();

  return `${d}.${m}.${y}`;
}

export function formatDateTime(date): string {
  date = assertDate(date);
  const h = date.getHours();
  const m = date.getMinutes();

  return `${formatDate(date)} ${h}:${m}`;
}
