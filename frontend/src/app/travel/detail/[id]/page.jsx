"use client";

import { useState, useEffect } from "react";
import { useParams } from "next/navigation";
import Link from "next/link";
import axios from "axios";

export default function Detail() {
    const { id } = useParams();

    const [detailTravelData, setDetailTravelData] = useState({});

    useEffect(() => {
        fetchData()
    }, [id])

    const fetchData = async () => {
        try {
            const detail = await axios(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${id}`)
            setDetailTravelData(detail.data.Data)
        } catch (error) {
            console.log("Terjadi error saat mengambil data detail:", error);
        }
    }

    function formatRupiah(angka) {
        return new Intl.NumberFormat('id-ID', {
            style: 'currency',
            currency: 'IDR',
            minimumFractionDigits: 0
        }).format(angka);
    }

    return (
        <div className="min-h-screen flex items-center justify-center px-8">
            <div className="bg-violet-300 shadow-lg rounded-lg p-6 max-w-4xl w-auto">

                {/* Judul di atas */}
                <h1 className="text-3xl font-bold mb-6 text-slate-700">
                    {detailTravelData.name}
                </h1>

                {/* Baris: Gambar + Detail */}
                <div className="flex flex-col md:flex-row items-start gap-8">
                    <img
                        src={detailTravelData.photo}
                        alt={detailTravelData.name}
                        className="w-80 h-80 object-cover rounded-md shadow-lg"
                    />

                    <div className="flex-1 text-start text-black">
                        <ul className="list-none space-y-4">
                            <li>
                                <span className="font-semibold mr-4">Harga:</span>
                                {formatRupiah(detailTravelData.price)},-
                            </li>
                            <li className="text-justify">
                                <span className="font-semibold mr-2">Deskripsi:</span>
                                {detailTravelData.description}
                            </li>
                        </ul>

                        <div className="flex justify-end mt-6">
                            <Link href="/">
                                <button className="bg-gray-400 hover:bg-gray-600 text-white font-semibold py-2 px-5 rounded-lg shadow-md">
                                    Kembali
                                </button>
                            </Link>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}