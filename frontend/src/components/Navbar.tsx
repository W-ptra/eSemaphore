function Navbar() {
    return (
        <nav
            className="w-15 bg-white flex flex-col items-center py-2 z-20"
            style={{ boxShadow: '2px 0 4px rgba(0, 0, 0, 0.1)' }}
        >
            <ul
                className="flex-1/2 flex flex-col gap-y-3"
            >
                <li className="hover:bg-gray-100 p-1 rounded-xl cursor-pointer" >

                    <img
                        src="https://img.icons8.com/sf-black-filled/64/228BE6/chat-message.png" alt=""
                        className="size-9"
                    />
                </li>
            </ul>
            <ul
                className="flex-1/2 justify-end flex flex-col gap-y-3"
            >
                <li 
                    className="hover:bg-gray-100 p-1 rounded-xl cursor-pointer"
                >
                    <img
                        src="https://img.icons8.com/fluency-systems-filled/50/228BE6/user-male-circle.png" alt=""
                        className="size-9"
                    />
                </li>
            </ul>
        </nav>
    )
}

export default Navbar