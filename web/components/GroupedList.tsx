import React, { ReactNode } from 'react';

interface Props<T> {
  groupItems: (items: T[]) => T[][];
  renderHeader: (item: T[]) => ReactNode;
  renderList: (items: T[]) => ReactNode;
  items: T[];
}

export default function GroupedList<T>({
  groupItems,
  items,
  renderHeader,
  renderList,
}: Props<T>) {
  const groupedItems = groupItems(items);

  const groupsWithHeaders = [];

  groupedItems.forEach(group => {
    groupsWithHeaders.push(renderHeader(group));
    groupsWithHeaders.push(renderList(group));
  });

  return <>{groupsWithHeaders}</>;
}
