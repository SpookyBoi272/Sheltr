import { useEffect, useRef, useState } from "react";
import { Link, useNavigate } from "react-router";
import axios from "../api/axios";
import { useMutation } from "react-query";
import { IconButton, Input } from "@mui/material";
import { Visibility, VisibilityOff } from "@mui/icons-material";

const NAME_REGEX = /^[a-z ,.'-]{3,50}$/i;
const EMAIL_REGEX = /^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$/;
const PWD_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%]).{8,24}$/;

const input_style = {
  "&.MuiInput-root": {
    padding: "0px 30px 0px 40px",
  },
  "& input": {
    color: "#49735A",
    padding: "0px",
    fontSize: "18px",
    fontFamily: "",
  },
  "& input::placeholder": {
    opacity: 1,
  },
  "&:focus-within": {
    borderColor: "#49735A",
    outline: "2px solid #97b1a1",
  },
};

const REGISTER_URL = "/fact";

export default function Register() {
  const userRef = useRef();
  const errorRef = useRef();

  const [name, setName] = useState("");
  const [validName, setValidName] = useState(false);

  const [email, setEmail] = useState("");
  const [validEmail, setValidEmail] = useState(false);

  const [password, setPassword] = useState("");
  const [validPassword, setValidPassword] = useState(false);

  const [matchPassword, setMatchPassword] = useState("");
  const [validMatch, setValidMatch] = useState(false);

  const [showPassword, setShowPassword] = useState(false);
  const [showMatchPassword, setShowMatchPassword] = useState(false);

  const [errorMessage, setErrorMessage] = useState("");

  const navigate = useNavigate();

  useEffect(() => {
    userRef.current.focus();
  }, []);

  useEffect(() => {
    setValidName(NAME_REGEX.test(name));
  }, [name]);

  useEffect(() => {
    setValidEmail(EMAIL_REGEX.test(email));
  }, [email]);

  useEffect(() => {
    setValidPassword(PWD_REGEX.test(password));
    setValidMatch(password === matchPassword);
  }, [password, matchPassword]);

  const { mutateAsync: registerUser } = useMutation({
    mutationFn: (newUser) => {
      return axios.post(REGISTER_URL, newUser);
    },
    retry: 2,
  });

  const handleSubmit = async (e) => {
    e.preventDefault();
    const v1 = EMAIL_REGEX.test(email);
    const v2 = PWD_REGEX.test(password);
    const v3 = NAME_REGEX.test(name);
    const v4 = password === matchPassword;
    if (!v1 || !v2 || !v3 || !v4) {
      setErrorMessage(" Invalid Entry");
      return;
    }
    try {
      const response = await registerUser({ name, email, password });
      console.log(response?.data);
      console.log(response?.accessToken);
      setName("");
      setEmail("");
      setPassword("");
      setMatchPassword("");
      navigate("/");
    } catch (err) {
      if (!err?.response) {
        setErrorMessage("No Server Response");
      } else if (err.response?.status === 409) {
        setErrorMessage("Email already in use");
      } else {
        setErrorMessage("Registration Failed, Please try again later.");
      }
      errorRef.current.focus();
    }
  };

  return (
    <section>
      <div className="w-[510px] mx-auto p-[40px] flex flex-col mt-12 rounded-3xl">
        <p
          ref={errorRef}
          className={errorMessage ? "errmsg" : "offscreen"}
          aria-live="assertive"
        >
          {errorMessage}
        </p>
        <p className="heading mx-auto">Create your account</p>
        <form
          onSubmit={handleSubmit}
          className="flex flex-col my-5 register-form"
        >
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
          <Input
            type={showPassword ? "text" : "password"}
            name="password"
            id="password"
            className="input"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            disableUnderline={true}
            sx={input_style}
            aria-describedby="password-note"
            aria-invalid={validPassword ? "false" : "true"}
            endAdornment={
              <IconButton onClick={() => setShowPassword((m) => !m)} edge="end">
                {!showPassword ? <Visibility /> : <VisibilityOff />}
              </IconButton>
            }
          />
          <p
            id="password-note"
            className={password && !validPassword ? "requirements" : "hidden"}
          >
            Must be 8 to 24 characters.
            <br />
            Must include at least one uppercase letter, one lowercase letter,
            one number, and one special character (!@#$%).
          </p>
          <Input
            type={showMatchPassword ? "text" : "password"}
            name="match"
            id="match"
            className="input"
            value={matchPassword}
            onChange={(e) => setMatchPassword(e.target.value)}
            placeholder="Confirm Password"
            disableUnderline={true}
            sx={input_style}
            required
            aria-describedby="confirm-password-note"
            aria-invalid={!validMatch ? "false" : "true"}
            endAdornment={
              <IconButton
                onClick={() => setShowMatchPassword((m) => !m)}
                edge="end"
              >
                {showMatchPassword ? <Visibility /> : <VisibilityOff />}
              </IconButton>
            }
          />
          <p
            id="confirm-password-note"
            className={matchPassword && !validMatch ? "requirements" : "hidden"}
          >
            Must match the password
          </p>
          <button
            className="bg-[#064749] text-white font-bold h-[48px] w-[180px] rounded-full mt-4 mx-auto"
            // disabled={
            //   !validName || !validPassword || !validEmail || !validMatch
            //     ? true
            //     : false
            // }
          >
            Create Account
          </button>
        </form>
        <p>
          Already registered?{" "}
          <span>
            <Link to="/login">Sign In</Link>
          </span>
        </p>
      </div>
    </section>
  );
}
