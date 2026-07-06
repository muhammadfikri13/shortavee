import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import API from "../services/api";
import toast from "react-hot-toast";
import Swal from 'sweetalert2';


export default function Dashboard() {
    const [urls, setUrls] = useState([]);
    const navigate = useNavigate();
    const [newURL, setNewURL] = useState("");
    const [loading, setLoading] = useState(false);
    const [fetching, setFetching] = useState(true);
    const VITE_API_BASE_URL = import.meta.env.VITE_API_BASE_URL || "http://localhost:8080";

    // 1. Buat state baru untuk memicu proses pengambilan data ulang
    const [refreshTrigger, setRefreshTrigger] = useState(0);

    useEffect(() => {
        // 2. Kembalikan fetchUrls ke dalam useEffect agar linter bahagia
        const fetchUrls = async () => {

            setFetching(true);

            try {
                const res = await API.get("/api/urls");
                console.log(res.data);
                setUrls(res.data);
            } catch (err) {
                console.error(err);
            } finally {
                setFetching(false);
            }
        };

        fetchUrls();
        
    // 3. Masukkan refreshTrigger ke array dependensi.
    // Artinya: Setiap kali 'refreshTrigger' berubah nilainya, useEffect ini akan otomatis dijalankan ulang.
    }, [refreshTrigger]); 

    const createURL = async (e) => {
        e.preventDefault();

        setLoading(true);
        try {
            await API.post("/api/shorten", {
                url: newURL,
            });

            setNewURL("");
            toast.success("URL created!")
            
            // 4. Ubah nilai trigger (tambah 1) untuk memberi tahu useEffect agar mengambil data terbaru
            setRefreshTrigger((prev) => prev + 1);
        }   catch (err) {
            console.error(err);
            toast.error("Failed to create URL");
        }   finally {
            setLoading(false);
        }
    };

    const deleteURL = async (id) => {
        // Memanggil SweetAlert2, ini akan mengembalikan Promise
        const result = await Swal.fire({
            title: "Delete this URL?",
            text: "You won't be able to revert this!",
            icon: "warning",
            showCancelButton: true,
            confirmButtonColor: "#dc2626", // Warna merah Tailwind
            cancelButtonColor: "#9ca3af", // Warna abu-abu Tailwind
            confirmButtonText: "Yes, Delete"
        });

        // Jika user klik "Cancel" atau menutup popup, hentikan fungsi
        if (!result.isConfirmed) return;

        try {
            await API.delete(`/api/urls/${id}`);
            setUrls((prevUrls) => prevUrls.filter((url) => url.id !== id));
            Swal.fire({
                title: "Deleted!",
                text: "Your URL has been deleted successfully.",
                icon: "success",
                timer: 1500,
                showConfirmButton: false
            });
        } catch (err) {
            console.error(err);
            toast.error("Failed to delete URL");
        }
    }

    const copyToClipboard = (shortCode) => {

        const shortURL =
            `${VITE_API_BASE_URL}/${shortCode}`;
        
            navigator.clipboard.writeText(shortURL);
            toast.success("Copied!");
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
                className="flex bg-white p-4 rounded-lg shadow-mb-6">
                    <input
                        type="text"
                        placeholder="https://example.com"
                        value={newURL}
                        onChange={(e) => setNewURL(e.target.value)}
                        className="flex-1 border border-gray-300 rounded px-3 py-2 mr-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />

                    <button 
                        type="submit"
                        disabled={loading}
                        className="bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded"
                    >
                        {loading ? "Creating..." : "Shorten"}
                    </button>
                </form>
            </div>
            
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 p-6">
                {urls.map((url) => (
                    <div
                    key={url.id}
                    className="bg-white rounded-lg shadow p-4"
                    >

                    <h3 className="font-semibold text-gray-800">
                        Short URL
                    </h3>

                    <p className="text-blue-600 mb-3">
                        {VITE_API_BASE_URL}/{url.short_code}
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

                        <div className="flex items-center gap-2">
                            <button
                                onClick={() =>
                                    copyToClipboard(url.short_code)
                                }
                                className="bg-indigo-600 hover:bg-indigo-700 text-white px-3 py-1 rounded"
                                >
                                Copy
                            </button>

                            <button
                                onClick={() =>
                                deleteURL(url.id)
                                }
                                className="bg-white border-2 border-red-200 hover:bg-red-600 text-black hover:text-white px-3 py-1 rounded"
                                >
                                Delete
                            </button>
                        </div>
                    </div>


                    </div>
                ))}
            </div>
        </div>
    )
}