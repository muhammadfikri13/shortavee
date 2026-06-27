import { useState } from "react";
import API from "../services/api";
import { useNavigate } from "react-router-dom";    

export default function Login() {

    const navigate = useNavigate();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleLogin = async (e) => {
    e.preventDefault();

    try {
        const res = await API.post("/api/login", {
            email,
            password,
        });

        localStorage.setItem("token", res.data.token);

        navigate("/dashboard");
      } catch {
        alert("Invalid email or password");
    }
};

    return (
        <div>
            <h1 className="text-4xl font-bold">
                Login
            </h1>
            <form onSubmit={handleLogin}>
                <input
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />

                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />

                <button type="submit">
                    Login
                </button>
            </form>
        </div>
    )
}