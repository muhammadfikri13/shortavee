import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import API from "../services/api";

export default function Dashboard() {
    const [urls, setUrls] = useState([]);
    const navigate = useNavigate();
    const [newURL, setNewURL] = useState("");
    
    // 1. Buat state baru untuk memicu proses pengambilan data ulang
    const [refreshTrigger, setRefreshTrigger] = useState(0);

    useEffect(() => {
        // 2. Kembalikan fetchUrls ke dalam useEffect agar linter bahagia
        const fetchUrls = async () => {
            try {
                const res = await API.get("/api/urls");
                console.log(res.data);
                setUrls(res.data);
            } catch (err) {
                console.error(err);
            }
        };

        fetchUrls();
        
    // 3. Masukkan refreshTrigger ke array dependensi.
    // Artinya: Setiap kali 'refreshTrigger' berubah nilainya, useEffect ini akan otomatis dijalankan ulang.
    }, [refreshTrigger]); 

    const createURL = async (e) => {
        e.preventDefault();

        try {
            await API.post("/api/shorten", {
                url: newURL,
            });

            setNewURL("");
            
            // 4. Ubah nilai trigger (tambah 1) untuk memberi tahu useEffect agar mengambil data terbaru
            setRefreshTrigger((prev) => prev + 1);
        } catch (err) {
            console.error(err);
            alert("Failed to create URL");
        }
    };

    const deleteURL = async (id) => {

        const confirmed = window.confirm(
            "Delete this URL?"
        );

        if (!confirmed) return;

        try {
            await API.delete(`/api/urls/${id}`);

        }   catch (err) {
            console.error(err);
        }
    }

    const copyToClipboard = (shortCode) => {

        const shortURL =
            `http://localhost:8080/${shortCode}`;
        
            navigator.clipboard.writeText(shortURL);

            alert("Copied!");
    };

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/");
    }

    return (
        <div className="min-h-screen bg-gray-100">
            {/* Header */}
            <div className="bg-white shadow">
                <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
                    <h1 className="text-2xl font-bold">
                        MyURLs
                    </h1>
                    <button 
                        onClick={handleLogout}
                        className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded absolute top-0 right-0 m-5"
                    >Logout
                    </button>
                </div>
            </div>

            <div className="max-w-5xl mx-auto p-6">
                <form 
                onSubmit={createURL}
                className="bg-white p-4 rounded-lg shadow-mb-6">
                    <input
                        type="text"
                        placeholder="https://example.com"
                        value={newURL}
                        onChange={(e) => setNewURL(e.target.value)}
                        className="flex-1 border border-gray-300 rounded px-3 py-2 mr-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />

                    <button 
                    type="submit"
                    className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                    >
                        Shorten
                    </button>
                </form>

            </div>
            
            <div className="grid gap-4">
                {urls.map((url) => (
                    <div
                    key={url.id}
                    className="bg-white rounded-lg shadow p-4"
                    >

                    <h3 className="font-semibold text-gray-800">
                        Short URL
                    </h3>

                    <p className="text-blue-600 mb-3">
                        http://localhost:8080/{url.short_code}
                    </p>

                    <h3 className="font-semibold text-gray-800">
                        Original URL
                    </h3>

                    <p className="text-gray-600 break-all mb-3">
                        {url.original_url}
                    </p>

                    <div className="flex justify-between items-center">

                        <span className="text-sm text-gray-500">
                        Clicks: {url.click_count}
                        </span>

                        <button
                        onClick={() =>
                            copyToClipboard(url.short_code)
                        }
                        className="bg-green-600 text-white px-3 py-1 rounded"
                        >
                        Copy
                        </button>

                        <button
                            onClick={() =>
                                deleteURL(url.id)
                            }
                            className="bg-red-600 text-white px-3 py-1 rounded"
                            >
                                Delete
                        </button>

                    </div>

                    </div>
                ))}
            </div>
        </div>
    )
}