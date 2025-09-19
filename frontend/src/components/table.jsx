"use client";
import { useEffect, useState } from "react";
import Link from "next/link";
import axios from "axios";
import Swal from "sweetalert2";

export default function Table() {
    const [TravelData, setTravelData] = useState([]);

    useEffect(() => {
        fetchData();
    }, []);

    const fetchData = async () => {
        try {
            const get = await axios.get(process.env.NEXT_PUBLIC_BACKEND_URL);
            setTravelData(get.data.Data);
        } catch (error) {
            if (error.response) {
                console.log("Terjadi Kesalahan Saat Mengambil Data (Response):", error.response.status, error.response.data);
            } else if (error.request) {
                console.log("Terjadi Kesalahan Saat Mengambil Data (Request):", error.message);
            } else {
                console.log("Terjadi Kesalahan Saat Mengambil Data (General):", error.message);
            }
        }
    }

    function formatRupiah(angka) {
        return new Intl.NumberFormat('id-ID', {
            style: 'currency',
            currency: 'IDR',
            minimumFractionDigits: 0
        }).format(angka);
    }

    const handleDelete = async (id) => {
        Swal.fire({
            title: "Apakah Anda Yakin Ingin Menghapus Data Ini?",
            text: "Data yang sudah dihapus tidak dapat dikembalikan!",
            icon: "warning",
            showCancelButton: true,
            showConfirmButton: true,
            cancelButtonText: "Batal",
            confirmButtonText: "Hapus",
            confirmButtonColor: "#d33",
            cancelButtonColor: "#3085d6",
        }).then(async (result) => {
            if (result.isConfirmed) {
                try {
                    await axios.delete(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${id}`);

                    const newData = TravelData.filter((item) => {
                        return item.id !== id;
                    });
                    setTravelData(newData);

                    Swal.fire({
                        icon: "success",
                        title: "Berhasil Menghapus Data",
                        text: "Data berhasil dihapus dari database",
                    })
                } catch (error) {
                    Swal.fire({
                        icon: "error",
                        title: "Gagal Menghapus Data",
                        text: "Terjadi kesalahan saat menghapus data",
                    })
                }
            }
        })
    }

    return (
        <table>
            <thead className="text-bold text-zinc-800 uppercase text-center bg-lime-300">
                <tr>
                    <th className="py-4 px-7 text-center">No</th>
                    <th className="py-4 px-7 text-center">Name</th>
                    <th className="py-4 px-7 text-center">Price</th>
                    <th className="py-4 px-7 text-center">Image</th>
                    <th className="py-4 px-7 text-center">Actions</th>
                </tr>
            </thead>
            <tbody>
                {TravelData.map((item, index) => (
                    <tr className="bg-yellow-100 hover:bg-amber-200 text-slate-700" key={item.id}>
                        <td className="py-4 px-7 text-center">{index + 1}</td>
                        <td className="py-4 px-7 text-center">{item.name}</td>
                        <td className="py-4 px-7 text-center">{formatRupiah(item.price)},-</td>
                        <td className="py-4 px-7 text-center">
                            <img src={item.photo} alt={item.name} className="w-24 h-24 object-cover mx-auto" />
                        </td>
                        <td className="py-4 px-7 text-center flex gap-2 justify-betweens">
                            <Link href={`/travel/detail/${item.id}`}>
                                <button className="bg-cyan-400 hover:bg-cyan-600 text-white font-sm rounded-lg shadow-2xl p-3"> Details </button>
                            </Link>
                            <Link href={`/travel/update/${item.id}`}>
                                <button className="bg-indigo-400 hover:bg-indigo-600 text-white font-sm rounded-lg shadow-2xl p-3"> Update </button>
                            </Link>
                            <button className="bg-red-400 hover:bg-red-600 text-white font-sm rounded-lg shadow-2xl p-3" onClick={() => handleDelete(item.id)}> Delete </button>
                        </td>
                    </tr>
                ))}
            </tbody>
        </table>
    )
}