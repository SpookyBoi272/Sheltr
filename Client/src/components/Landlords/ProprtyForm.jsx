import { useEffect, useRef, useState } from "react";

const NAME_REGEX = /^[a-z ,.'-]{3,50}$/i;
const EMAIL_REGEX = /^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$/;

export default function ProprtyForm() {
  const userRef = useRef();
  const errorRef = useRef();

  const [name, setName] = useState("");
  const [validName, setValidName] = useState(false);

  const [email, setEmail] = useState("");
  const [validEmail, setValidEmail] = useState(false);

  const [errorMessage, setErrorMessage] = useState("");

  useEffect(() => {
    setValidName(NAME_REGEX.test(name));
  }, [name]);

  useEffect(() => {
    setValidEmail(EMAIL_REGEX.test(email));
  }, [email]);
  return (
    <div
      className="main"
      style={{
        background: `url(/landlord.png)`,
        backgroundRepeat: "no-repeat",
        backgroundSize: "cover",
        backgroundPosition: "left",
      }}
    >
      <section className="proprty-form">
        <p
          ref={errorRef}
          className={errorMessage ? "errmsg" : "offscreen"}
          aria-live="assertive"
        >
          {errorMessage}
        </p>
        <p className="heading mx-auto">Create your account</p>
        <p className="heading">Earn more from your property, do less</p>
        <p>Find out if your proprty meets our criteria</p>
        <form>
          <input
            type="text"
            name="name"
            id="name"
            className="input"
            onChange={(e) => {
              setName(e.target.value);
            }}
            placeholder="Name"
            ref={userRef}
            value={name}
            required
            aria-describedby="name-note"
            aria-invalid={validName ? "false" : "true"}
          />
          <p
            id="name-note"
            className={name && !validName ? "requirements" : "hidden"}
          >
            3 to 50 characters.
            <br />
            Can include letters, spaces, commas, periods, apostrophes, and
            hyphens.
            <br />
            Case-insensitive (uppercase or lowercase allowed).
          </p>
          <input
            type="text"
            name="email"
            id="email"
            className="input"
            onChange={(e) => {
              setEmail(e.target.value);
            }}
            placeholder="Email"
            value={email}
            required
            aria-describedby="email-note"
            aria-invalid={validEmail ? "false" : "true"}
          />
          <p
            id="email-note"
            className={email && !validEmail ? "requirements" : "hidden"}
          >
            Must have a domain name after the @ symbol, which can include
            lowercase letters (a-z), numbers (0-9), periods (.), and hyphens
            (-). <br />
            Must end with a valid top-level domain (e.g., .com, .org, .net,
            etc.) with at least 2 characters.{" "}
          </p>
        </form>
      </section>
    </div>
  );
}
