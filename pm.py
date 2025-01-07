def PontoMedio(self, x0, y0, xEnd, yEnd):
        dx = abs(xEnd - x0)
        dy = abs(yEnd - y0)

        if dx >= dy and x0 <= xEnd and y0 <= yEnd:# Verifica se a linha está no primeiro oitante
            print("1 oitante")
            ds = 2 * dy - dx
            incE = 2 * dy
            incNE = 2 * (dy - dx)
            x, y = 0, 0
            if x0 > xEnd:
                x, y = xEnd, yEnd
                xEnd = x0
            else:
                x, y = x0, y0
                
            self.points.append((round(x), round(y)))
            
            while x < xEnd:
                x += 1
                if ds < 0:
                    ds += incE
                else:
                    y += 1
                    ds += incNE
                self.points.append((round(x), round(y)))
            self.update()
            return 
        elif dx >= dy and x0 >= xEnd and y0 <= yEnd: # 4 oitante
            print("4 oitante")
            ds = 2 * dy - dx
            incE = 2 * dy
            incNE = 2 * (dy - dx)
            x, y = x0, y0
                
            self.points.append((round(x), round(y)))
            
            while x > xEnd:
                x -= 1
                if ds < 0:
                    ds += incE
                else:
                    y += 1
                    ds += incNE
                self.points.append((round(x), round(y)))
            self.update()
            return 
        elif dy >= dx and x0 <= xEnd and y0 <= yEnd: # 2 oitante 
            print("2 oitante")
            ds = 2 * dx - dy  # Ajustado para (y, x)
            incE = 2 * dx     # Ajustado para incremento E
            incNE = 2 * (dx - dy)  # Ajustado para incremento NE
            x, y = x0, y0
                
            self.points.append((round(x), round(y)))
            
            while y < yEnd:
                y += 1
                if ds < 0:
                    ds += incE
                else:
                    x += 1
                    ds += incNE
                self.points.append((round(x), round(y)))
            self.update()
            return

        elif (dy >= dx) and x0 >= xEnd and y0 <= yEnd: # 3 oitante
            print("3 oitante")
            ds = 2 * dx - dy  # Ajustado para (y, x)
            incE = 2 * dx     # Ajustado para incremento E
            incNE = 2 * (dx - dy)  # Ajustado para incremento NE
            x, y = x0, y0
                
            self.points.append((round(x), round(y)))
            
            while x > xEnd:
                y += 1
                if ds < 0:
                    ds += incE
                else:
                    x -= 1
                    ds += incNE
                self.points.append((round(x), round(y)))
            self.update()
            return 
        elif(dx>=dy) and x0 >= xEnd and y0 >= yEnd: # 5 oitante
            print("5 oitante")
            ds = 2 * -dy + dx
            incE = 2 * -dy
            incNE = 2 * (-dy + dx) # Ajustado para incremento NE
            x, y = x0, y0

            self.points.append((round(x), round(y)))


            while x > xEnd:  
                x -= 1
                if ds < 0:
                    ds += -incE
                else:
                    y -= 1
                    ds += -incNE
                self.points.append((round(x), round(y)))

            self.update()
            return
        elif(dy>=dx) and x0 >= xEnd and y0 >= yEnd: # 6 oitante
            print("6 oitante")
            dx, dy =  dy, dx
            ds = 2 * -dy + dx
            incE = 2 * -dy
            incNE = 2 * (-dy + dx) # Ajustado para incremento NE
            x, y = x0, y0

            self.points.append((round(x), round(y)))


            while y > yEnd:
                y -= 1
                if ds < 0:
                    ds += -incE
                else:
                    x -= 1
                    ds += -incNE
                self.points.append((round(x), round(y)))

            self.update()
            return
        elif(dy>=dx) and xEnd >= x0 and y0 >= yEnd: # 7 oitante
            print("7 oitante")
            dx, dy =  dy, dx
            ds = 2 * -dy + dx
            incE = 2 * -dy
            incNE = 2 * (-dy + dx) # Ajustado para incremento NE
            x, y = x0, y0

            self.points.append((round(x), round(y)))


            while y > yEnd:
                y -= 1
                if ds < 0:
                    ds += -incE
                else:
                    x += 1
                    ds += -incNE
                self.points.append((round(x), round(y)))

            self.update()
            return
        elif(dx>=dy) and x0 <= xEnd and y0 >= yEnd: # 8 oitante
            print("8 oitante")
            ds = 2 * -dy + dx
            incE = 2 * -dy
            incNE = 2 * (-dy + dx) # Ajustado para incremento NE
            x, y = x0, y0

            self.points.append((round(x), round(y)))


            while x < xEnd:
                x += 1
                print(ds, incE, incNE)
                if ds < 0:
                    ds += -incE
                else:
                    y -= 1
                    ds += -incNE
                self.points.append((round(x), round(y)))

            self.update()
            return