import { useState } from "react";
import BackHeader from "../components/BackHeader";

function Profile() {
    const [name,setName] = useState("")
    const [email,setEmail] = useState("")


    const nameChangehandler = (event: React.ChangeEvent<HTMLInputElement>) => {
        setName(event.target.value)
    }

    const emailChangehandler = (event: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(event.target.value)
    }

    return (
        <>

            <BackHeader/>

            <main
                className="flex justify-center mt-4 md:mt-10"
            >
                <div
                    className="w-[40rem] h-[28rem] shadow border border-gray-100 rounded-xl flex flex-col gap-y-5 justify-center"
                >
                    <div
                        className="flex justify-center items-center"
                    >
                        <img src="/logo/account_logo.png" alt="" 
                            className="size-28 bg-gray-100 rounded-full"
                        />
                    </div>
                    <div
                        className="flex flex-col items-center"
                    >
                        <div>
                            <label htmlFor="id" className="pl-2">ID</label>
                            <br />
                            <input type="text" value={"Id"}
                                className="border-2 border-gray-100 px-2 rounded" 
                                disabled={true}
                            />
                        </div>
                    </div>
                    <div
                        className="flex flex-col items-center"
                    >
                        <div>
                            <label htmlFor="id" className="pl-2">Name</label>
                            <br />
                            <input type="text" value={name || "name"}
                                className="border-2 border-gray-100 px-2 rounded" 
                                onChange={nameChangehandler}
                            />
                        </div>
                    </div>
                    <div
                        className="flex flex-col items-center"
                    >
                        <div>
                            <label htmlFor="id" className="pl-2">Email</label>
                            <br />
                            <input type="text" value={email || "email"}
                                className="border-2 border-gray-100 px-2 rounded" 
                                onChange={emailChangehandler}
                            />
                        </div>
                    </div>
                    <div
                        className="flex justify-center items-center gap-x-2"
                    >
                        <button
                            className="bg-blue-500 text-white font-semibold px-2 py-1 rounded hover:bg-blue-400 cursor-pointer"
                        >
                            Update Profile
                        </button>
                        <button
                            className="bg-blue-500 text-white font-semibold px-2 py-1 rounded hover:bg-blue-400 cursor-pointer"
                        >
                            Change Password
                        </button>
                    </div>
                </div>
            </main>
        </>
    )
}

export default Profile