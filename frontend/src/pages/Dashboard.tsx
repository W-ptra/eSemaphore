import Header from "../components/Header"
import Navbar from "../components/Navbar"
import ChatBar from "../components/ChatBar"
import ChatLine from "../components/ChatLine"

function Dashboard() {
    const lorem = "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Rerum expedita at, repudiandae debitis accusamus animi numquam. Ut, quaerat itaque omnis sequi libero sint voluptas, maiores excepturi mollitia porro ipsum quibusdam?"
    return (
        <div
            className=" fixed flex flex-col top-0 bottom-0 left-0 right-0 h-screen"
        >
            <Header />
            <div className="flex flex-1 h-[80%] flex-col-reverse md:flex-row">
                <Navbar />
                <div
                    className="bg-white flex-1 flex h-full"

                >
                    <div
                        className="hidden md:block flex-[0_0_30%] overflow-y-auto h-full custom-scrollbar border-2 border-gray-100"
                    >
                        <ChatBar />
                        <ChatBar />
                        <ChatBar />
                        <ChatBar />
                        <ChatBar />
                        <ChatBar />
                        <ChatBar />
                        <ChatBar />
                    </div>
                    <div
                        className="flex-[0_0_100%] md:flex-[0_0_70%] mt-1 flex flex-col h-full w-full"
                    >
                        <div
                            className="flex-[0_0_10%] border-2 border-white border-b-gray-100 px-2 flex items-center gap-x-2"
                        >
                            <img src="/logo/account_logo.png" alt="" 
                                className="size-9"
                            />
                            <h2>
                                Juan
                            </h2>
                        </div>
                        <div
                            className="overflow-y-auto custom-scrollbar flex-[0_0_80%] px-5 py-1"
                        >
                            <ChatLine message={lorem} isSender={true} />
                            <ChatLine message={lorem} isSender={false} />
                            <ChatLine message={lorem} isSender={false} />
                            <ChatLine message={"aaa"} isSender={true} />
                            <ChatLine message={lorem} isSender={true} />
                            <ChatLine message={"asndansda"} isSender={false} />
                        </div>
                        <div
                            className="flex-[0_0_10%] border-2 border-white border-t-gray-100 flex justify-center items-center px-2"
                        >
                            <div
                                className="flex-11/12"
                            >
                                <input
                                    type="text"
                                    className="w-full border border-white focus:outline-none focus:ring-0 focus:border-transparent rounded px-3 py-2"
                                    placeholder="Type a message..."
                                />

                            </div>

                            <div 
                                className="flex-1/12 flex justify-center items-center"
                            >
                                <div
                                    className="hover:bg-gray-200 p-1 rounded cursor-pointer"
                                >
                                    <img src="/logo/send_logo.png" alt="" 
                                        className="size-6 "
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Dashboard