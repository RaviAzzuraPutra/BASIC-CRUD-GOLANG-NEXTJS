import Table from "@/components/table";
import Link from "next/link";

export default function Home() {
  return (
    <div className="w-screen py-24 flex justify-center flex-col items-center px-6">
      <div className="flex items-center justify-between gap-3">
        <h1 className="text-4xl font-asimovian font-extrabold">
          BASIC CRUD GOLANG AND NEXT.JS
        </h1>
      </div>
      <h2 className="text-2xl font-doto mb-6">
        BY RAVI AZZURA PUTRA
      </h2>
      <div className="overflow-x-auto mt-3">
        <div className="mb-7 w-full text-center">
          <Link href={"travel/create"}>
            <button className="bg-pink-500 hover:bg-pink-700 text-white font-extrabold p-3 rounded-lg shadow-lg font-libertinus font-2xl">
              + CREATE DATA
            </button>
          </Link>
        </div>
        <div>
          <Table />
        </div>
      </div>
    </div>
  );
}
