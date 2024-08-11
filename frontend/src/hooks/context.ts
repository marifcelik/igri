import { useContext } from "react";
import { UserContext } from "@/context/userContext";
import { ChatContext } from "@/context/chatContext"

export function useChat(){
	const context = useContext(ChatContext)
	if (!context) throw new Error('useChat must be used within a ChatProvider')
	return context
}

export function useUser() {
  const context = useContext(UserContext);
  if (!context) throw new Error('useUser must be used within a UserProvider');
  return context;
}