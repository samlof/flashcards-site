import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
import NavItem from "./NavItem";

interface Props {
  path: string;
  label: string;
}
const NavLink = ({ path, label }: Props) => {
  const router = useRouter();

  return (
    <Link href={path} passHref>
      <NavItem selected={router.pathname === path}>{label}</NavItem>
    </Link>
  );
};

export default NavLink;
