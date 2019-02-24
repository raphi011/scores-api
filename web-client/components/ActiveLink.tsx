import React, { Children } from 'react';

import Link, { LinkProps } from 'next/link';
import { SingletonRouter, withRouter } from 'next/router';

interface Props extends LinkProps {
  children: JSX.Element;
  activeClassName: string;
  href: string;
  router: SingletonRouter;
}

const ActiveLink = withRouter(
  ({ router, activeClassName, children, ...props }: Props) => (
    <Link {...props}>
      {React.cloneElement(Children.only(children), {
        className: router.pathname === props.href ? activeClassName : null,
      })}
    </Link>
  ),
);

export default ActiveLink;
