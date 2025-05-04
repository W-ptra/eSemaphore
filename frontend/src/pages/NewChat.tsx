import { useState } from "react";
import BackHeader from "../components/BackHeader"
import SearchAccount from "../components/SearchAccount"

function NewChat(){
    const [email,setName] = useState("");

    const emailChangeHandler = (event: React.ChangeEvent<HTMLInputElement>) => {
        setName(event.target.value)
    }

    return (
    <>
        <BackHeader/>
        <div
            className="flex justify-center mt-5"
        >
            <div>
                <div
                    className="flex"
                >
                    <input type="text" value={email}
                        placeholder="Email"
                        className="border-2 border-gray-100 px-2 py-1 w-[70vw] md:w-[40vw] rounded mr-1"
                        onChange={emailChangeHandler}
                    />
                    <button
                        className="bg-blue-500 text-white font-semibold px-2 py-1 rounded hover:bg-blue-400 cursor-pointer"
                    >
                        Search
                    </button>
                </div>
                <div>
                    <SearchAccount/>
                    <SearchAccount/>
                    <SearchAccount/>
                    <SearchAccount/>
                </div>
            </div>
        </div>
    </>
    )
}

export default NewChat