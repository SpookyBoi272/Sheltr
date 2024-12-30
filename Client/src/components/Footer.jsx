import Logo from "../assets/logo.svg";

export default function Footer() {
  return (
    <footer>
      <div className="flex flex-col w-1/3 gap-2">
        <img src={Logo} className="logo" alt="Rent Logo" />
        <span>Contact Number: 0000000000000</span>
        <ul className="flex gap-2">
          <li>
            <a href="https://linkedin.com">LinkedIn</a>
          </li>
          <li>
            <a href="https://facebook.com">Facebook</a>
          </li>
          <li>
            <a href="https://x.com">X</a>
          </li>
        </ul>
        <span>&copy; Rent</span>
      </div>
      <div className="flex w-1/3 gap-8 [&>ul]:flex [&>ul]:gap-2 [&>ul]:flex-col">
        <ul>
          <li className="font-bold">Company</li>
          <li>
            <a href="">Home</a>
          </li>
          <li>
            <a href="">About us</a>
          </li>
          <li>
            <a href="">Our team</a>
          </li>
        </ul>
        <ul>
          <li className="font-bold">Guests</li>
          <li>
            <a href="">Blog</a>
          </li>
          <li>
            <a href="">FAQ</a>
          </li>
          <li>
            <a href="">Career</a>
          </li>
        </ul>
        <ul>
          <li className="font-bold">Privacy</li>
          <li>
            <a href="">Terms of Servies</a>
          </li>
          <li>
            <a href="">Privacy Policy</a>
          </li>
        </ul>
      </div>
    </footer>
  );
}
