declare
fun {Prefix L1 L2}
   case L1
   of nil then true
   []H|T then
      case L2
      of nil then false
      []H2|T2 then
	 if H==H2 then {Prefix T T2}
	 else false end
      end
   end
end

{Browse {Prefix [1 2 5] [1 2 5]}}

declare
fun {FindString L1 L2}
   if {Prefix L1 L2} then true
   elseif L2==nil then false
   else {FindString L1 L2.2} end
end
      
   