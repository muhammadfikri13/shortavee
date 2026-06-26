import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import API from "../services/api";

export default function Dashboard() {

    const [urls, setUrls] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchUrls = async () => {
            try {
                const res = await API.get("/api/urls");

                console.log(res.data);
                setUrls(res.data);
            }   catch (err) {
                console.error(err);
            }
    };
        fetchUrls();
    }, []);

    

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/");
    }

    return (
        <div>
            <h1>
                My URLs
            </h1>

            {urls.map((url) => {
                return (
                    <div key={url.id}>
                        <p>Short URL: http://localhost:8080/{url.short_code}</p>
                        <p>Original URL: {url.original_url}</p>
                        <p>Clicks: {url.click_count}</p>
                        <br />
                    </div>
                );
            })}

            <button onClick={handleLogout}>
                Logout
            </button>
        </div>
    )
}
             