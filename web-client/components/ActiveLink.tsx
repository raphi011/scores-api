import React, { Children } from 'react';

import Link from 'next/link';
import { SingletonRouter, withRouter } from 'next/router';

interface Props {
  children: JSX.Element;
  activeClassName: string;
  href: string;
  router: SingletonRouter;
}

const ActiveLink = withRouter(
  ({ router, activeClassName, children, ...props }: Props) => (
    <Link {...props}>
      {React.cloneElement(Children.only(children), {
        className:
          `/${router.pathname.split('/')[1]}` === props.href
            ? activeClassName
            : null,
      })}
    </Link>
  ),
);

export default ActiveLink;
