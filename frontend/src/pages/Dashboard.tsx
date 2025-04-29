import Header from "../components/Header"
import Navbar from "../components/Navbar"

function Dashboard() {
    return (
        <div
            className=" fixed flex flex-col top-0 bottom-0 left-0 right-0 h-screen"
        >
            <Header />
            <div className="flex flex-1 h-[80%]">
                <Navbar />
                <div 
                    className="bg-white flex-1 flex h-full"
                    
                >
                    <div 
                        className="flex-1/4 overflow-y-auto h-full"
                    >
                        
                        <div className="h-14 shadow  m-5 rounded-xl flex items-center px-2">
                            <p>
                                chat
                            </p>
                        </div>
                        <div className="h-14 shadow  m-5 rounded-xl flex items-center px-2">
                            <p>
                                chat
                            </p>
                        </div>
                        <div className="h-14 shadow  m-5 rounded-xl flex items-center px-2">
                            <p>
                                chat
                            </p>
                        </div>
                        
                    </div>
                    <div
                        className="flex-3/4 bg-green-200"
                    >

                    </div>
                </div>
            </div>

            <div
                className="fixed bottom-0 right-0 m-5 bg-blue-500 size-20 flex items-center justify-center rounded-2xl hover:bg-blue-400 cursor-pointer"
            >
                <img
                    src="https://img.icons8.com/fluency-systems-filled/50/FFFFFF/add-to-chat.png" alt=""
                    className="size-12"
                />
            </div>
        </div>
    )
}

export default Dashboard