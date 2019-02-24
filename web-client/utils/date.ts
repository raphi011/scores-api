type DateType = Date | string;

function assertDate(date: DateType): Date {
  return date instanceof Date ? date : new Date(date);
}

export function formatDate(date: DateType): string {
  date = assertDate(date);
  const d = date.getDate();
  const m = date.getMonth() + 1;
  const y = date.getFullYear();

  return `${d}.${m}.${y}`;
}

export function formatDateTime(date: DateType): string {
  date = assertDate(date);
  const h = date.getHours();
  const m = date.getMinutes();

  return `${formatDate(date)} ${h}:${m}`;
}

export function sameDay(d1: Date, d2: Date): boolean {
  return (
    d1.getFullYear() === d2.getFullYear() &&
    d1.getMonth() === d2.getMonth() &&
    d1.getDay() === d2.getDay()
  );
}
