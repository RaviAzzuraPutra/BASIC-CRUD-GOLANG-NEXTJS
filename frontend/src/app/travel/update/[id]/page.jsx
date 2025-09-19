"use client"
import axios from "axios"
import Link from "next/link"
import { useParams } from "next/navigation"
import { useEffect, useState } from "react"
import Swal from "sweetalert2"

export default function Update() {
    const { id } = useParams();

    const [originalName, setOriginalName] = useState("")
    const [travelFields, setTravelFields] = useState({
        name: "",
        description: "",
        price: "",
        photo: "",
    })

    useEffect(() => {
        fetchData()
    }, [id])

    const fetchData = async () => {
        try {
            const get = await axios.get(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${id}`)
            setTravelFields(get.data.Data)
            setOriginalName(get.data.Data.name)
        } catch (error) {
            console.log("Terjadi Kesalahan Saat Mengambil Data:", error);
        }
    }

    const changeTravelFields = (e) => {
        setTravelFields({
            ...travelFields,
            [e.target.name]: e.target.value
        })
    }

    const onSubmit = async (e) => {
        e.preventDefault();

        const formData = new FormData();
        formData.append("name", travelFields.name)
        formData.append("description", travelFields.description)
        formData.append("price", travelFields.price)

        if (travelFields.photo && travelFields.photo instanceof File) {
            formData.append("photo", travelFields.photo)
        }

        try {
            const update = await axios.put(`${process.env.NEXT_PUBLIC_BACKEND_URL}/update-travel/${id}`, formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            });
            Swal.fire({
                icon: "success",
                title: "Berhasil Mengubah Data",
                text: update.data.Message
            }).then(() => {
                window.location.href = "/"
            })
        } catch (error) {
            Swal.fire({
                icon: "error",
                title: "Gagal Mengubah Data",
                text: error.response ? error.response.data.Message : error.message
            })
            console.log("Gagal Mengubah Data:", error);
        }
    }

    return (
        <div className="w-screen py-20 flex justify-center flex-col items-center px-7">
            <div className="flex items-center justify-between gap-3 mb-5">
                <h1 className="text-4xl font-semibold text-center"> Mengubah Data Travel "{originalName}" </h1>
            </div>
            <div className="overflow-x-auto w-full mt-4 bg-linear-65 from-cyan-500 to-blue-400 p-8 rounded-lg shadow-xl">
                <form action="">
                    <div className="mb-6">
                        <label htmlFor="title" className="block text-xl font-medium text-white text-bold mb-3">Nama</label>
                        <input
                            type="text"
                            name="name"
                            id="name"
                            placeholder="Masukan Nama Travel"
                            onChange={e => changeTravelFields(e)}
                            value={travelFields.name}
                            className="rounded-lg shadow-md w-full px-5 py-2 border border-green-400"
                        />
                    </div>
                    <div className="mb-6">
                        <label htmlFor="description" className="block text-xl font-medium text-white text-bold mb-3">Deskripsi</label>
                        <input
                            type="text"
                            name="description"
                            id="description"
                            placeholder="Masukan Deskripsi Travel"
                            onChange={e => changeTravelFields(e)}
                            value={travelFields.description}
                            className="rounded-lg shadow-md w-full px-5 py-2 border border-green-400"
                        />
                    </div>
                    <div className="mb-6">
                        <label htmlFor="price" className="block text-xl font-medium text-white text-bold mb-3">Harga</label>
                        <input
                            type="text"
                            name="price"
                            id="price"
                            placeholder="Masukan Harga Travel"
                            onChange={e => changeTravelFields(e)}
                            value={travelFields.price}
                            className="rounded-lg shadow-md w-full px-5 py-2 border border-green-400"
                        />
                    </div>
                    <div className="mb-6">
                        <label htmlFor="photo" className="block text-xl font-medium text-white text-bold mb-3">Photo</label>
                        <input
                            type="file"
                            name="photo"
                            id="photo"
                            placeholder="Masukan Photo Travel"
                            onChange={e => {
                                if (e.target.files && e.target.files.length > 0) {
                                    setTravelFields({
                                        ...travelFields,
                                        photo: e.target.files[0]
                                    })
                                }
                            }}
                            className="rounded-lg shadow-md w-full px-5 py-2 border border-green-400"
                        />
                    </div>

                    <div className="mb-6 flex justify-end mt-4 gap-4">
                        <button
                            type="submit"
                            onClick={e => onSubmit(e)}
                            className="bg-pink-400 hover:bg-pink-600 text-white font-semibold rounded-lg shadow-xl px-2 py-3">
                            SUBMIT
                        </button>
                        <Link href={"/"}>
                            <button className="bg-stone-400 hover:bg-stone-600 text-white font-semibold rounded-lg shadow-xl px-2 py-3">
                                KEMBALI
                            </button>
                        </Link>
                    </div>
                </form>
            </div>
        </div>
    )
}