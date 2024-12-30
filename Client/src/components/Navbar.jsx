import { NavLink } from "react-router";
import Logo from "../assets/logo.svg";
export default function Navbar() {
  return (
    <nav className="nav-bar">
      <div className="logo">
        <img src={Logo} className="logo" alt="Rent Logo" />
      </div>
      <div className="nav-items">
        {["Home", "Landlords", "Blog", "Contacts"].map((item) => (
          <NavLink
            to={"/" + item.toLocaleLowerCase()}
            key={item.toLocaleLowerCase()}
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            {item}
          </NavLink>
        ))}
      </div>
    </nav>
  );
}
