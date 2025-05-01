function ChatBar() {
    return (
        <div className="h-14 shadow  m-5 rounded-xl flex items-center px-2  cursor-pointer hover:bg-gray-100">
            <div
                className="basis-1/5 items-center  flex justify-center"
            >
                <img
                    src="https://img.icons8.com/sf-black-filled/64/228BE6/chat-message.png" alt=""
                    className="size-9 bg-gray-200 rounded-full"
                />
            </div>
            <div
                className="basis-3/5 flex flex-col justify-center text-sm"
            >
                <div>
                    <p>
                        chat name
                    </p>
                </div>
                <div>
                    <p>
                        123456789012345678901234
                    </p>
                </div>

            </div>
            <div
                className="basis-1/5 bg-amber-150 items-center  flex justify-center "
            >
                <p
                    className="text-end"
                >
                    11:42
                </p>
            </div>
        </div>
    )
}

export default ChatBar